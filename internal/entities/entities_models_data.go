package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Data struct {
	ID              primitive.ObjectID `bson:"_id"`
	SeriesReference string             `bson:"series_reference,omitempty"`
	Period          string             `bson:"period,omitempty"`
	DataValue       string             `bson:"data_value,omitempty"`
	Suppressed      string             `bson:"suppressed,omitempty"`
	Status          string             `bson:"status,omitempty"`
	Units           string             `bson:"units,omitempty"`
	Magnitude       string             `bson:"magnitude,omitempty"`
	Subject         string             `bson:"subject,omitempty"`
	Group           string             `bson:"group,omitempty"`
	SeriesTitle_1   string             `bson:"series_title_1,omitempty"`
	SeriesTitle_2   string             `bson:"series_title_2,omitempty"`
	SeriesTitle_3   string             `bson:"series_title_3,omitempty"`
	SeriesTitle_4   string             `bson:"series_title_4,omitempty"`
	SeriesTitle_5   string             `bson:"series_title_5,omitempty"`
}
