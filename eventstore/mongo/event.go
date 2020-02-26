package mongo

import (
	"encoding/json"

	"github.com/eiji03aero/mskit"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventDocument struct {
	Id            *primitive.ObjectID `bson:"_id,omitempty"`
	Type          string              `bson:"type"`
	AggregateType string              `bson:"aggregate_type"`
	AggregateId   string              `bson:"aggregate_id"`
	JsonData      []byte              `bson:"json_data"`
	Data          interface{}         `bson:"-"`
}

func NewEventDocument(event *mskit.Event) *EventDocument {
	return &EventDocument{
		Type:          event.Type,
		AggregateType: event.AggregateType,
		AggregateId:   event.AggregateId,
		Data:          event.Data,
	}
}

func (ed *EventDocument) makeJsonData() error {
	bytes, err := json.Marshal(ed.Data)
	if err != nil {
		return err
	}

	ed.JsonData = bytes
	return nil
}
