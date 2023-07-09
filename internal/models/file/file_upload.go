package file

import (
	"bufio"
	"context"
	"mime/multipart"
	"strings"
	"sync"

	"github.com/sadaghiani/concurrent-file-processing/internal/entities"
	"github.com/sadaghiani/concurrent-file-processing/internal/utils"
	"github.com/sadaghiani/concurrent-file-processing/pkg/config"
	"github.com/sadaghiani/concurrent-file-processing/pkg/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type dataBatch []interface{}

func (m *File) Upload(ctx context.Context, inputFile *multipart.FileHeader) error {

	numberWorkers := config.Config.GetInt(utils.MustBindEnvToString(utils.APP_NUMBER_WORKERS))

	openFile, err := inputFile.Open()
	if err != nil {
		logger.Log.Fatal("Unable to read input file "+inputFile.Filename, zap.Error(err))
		return err
	}
	defer openFile.Close()

	rowsBatch := []string{}
	rowsCh := read(ctx, &rowsBatch, openFile)

	workersCh := make([]<-chan dataBatch, numberWorkers)
	for i := 0; i < numberWorkers; i++ {
		workersCh[i] = process(ctx, rowsCh, i)
	}

	return m.Save(ctx, workersCh...)
}

func read(ctx context.Context, rowsBatch *[]string, openFile multipart.File) <-chan []string {

	out := make(chan []string)
	scanner := bufio.NewScanner(openFile)
	batchSize := config.Config.GetInt(utils.MustBindEnvToString(utils.APP_BATCH_SIZE))

	//skip title
	scanner.Scan()

	go func() {
		defer close(out)

		for {
			scanned := scanner.Scan()
			select {
			case <-ctx.Done():
				return
			default:
				row := scanner.Text()
				if len(*rowsBatch) == batchSize || !scanned {
					out <- *rowsBatch
					*rowsBatch = []string{}
				}
				*rowsBatch = append(*rowsBatch, row)
			}
			if !scanned {
				return
			}
		}
	}()

	return out
}

func process(ctx context.Context, rowBatch <-chan []string, wID int) <-chan dataBatch {

	out := make(chan dataBatch)

	go func() {
		defer close(out)

		d := dataBatch{}
		for rowBatch := range rowBatch {
			for _, row := range rowBatch {
				fields := strings.Split(row, ",")
				d = append(d, entities.Data{
					ID:              primitive.NewObjectID(),
					SeriesReference: fields[0],
					Period:          fields[1],
					DataValue:       fields[2],
					Suppressed:      fields[3],
					Status:          fields[4],
					Units:           fields[5],
					Magnitude:       fields[6],
					Subject:         fields[7],
					Group:           fields[8],
					SeriesTitle_1:   fields[9],
					SeriesTitle_2:   fields[10],
					SeriesTitle_3:   fields[11],
					SeriesTitle_4:   fields[12],
					SeriesTitle_5:   fields[13],
				})
			}
		}
		out <- d
	}()

	return out
}

func (f *File) Save(ctx context.Context, inputs ...<-chan dataBatch) error {

	var wg sync.WaitGroup
	saveInDatabase := func(p <-chan dataBatch) {
		defer wg.Done()

		for in := range p {
			select {
			case <-ctx.Done():
			default:
				_ = f.Repository.IRepositoryMongo.InsertBatch(ctx, in)
			}
		}
	}
	wg.Add(len(inputs))

	for _, in := range inputs {
		go saveInDatabase(in)
	}

	go func() {
		wg.Wait()
	}()

	return nil
}
