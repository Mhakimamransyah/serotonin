package pilgan

import (
	"context"
	"fmt"
	"serotonin/business/pilgan"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PilganRepository struct {
	collection string
	client     *mongo.Database
	ctx        *context.Context
}

func InitPilganRepository(client *mongo.Database, ctx *context.Context, collection string) *PilganRepository {
	return &PilganRepository{
		client:     client,
		ctx:        ctx,
		collection: collection,
	}
}

func (repository *PilganRepository) InsertDataPilgan(pilgan *pilgan.PilganTone) error {
	_, err := repository.client.Collection(repository.collection).InsertOne(*repository.ctx, *pilgan)
	if err != nil {
		return err
	}
	return nil
}

func (repository *PilganRepository) GetDataPilganById(id int) (*pilgan.PilganTone, error) {
	var pilgan pilgan.PilganTone
	filter := bson.M{"id_data_core": id}
	sort := bson.M{"version": -1}
	opts := options.FindOne().SetSort(sort)
	err := repository.client.Collection(repository.collection).FindOne(*repository.ctx, filter, opts).Decode(&pilgan)
	if err != nil {
		return nil, err
	}
	return &pilgan, nil
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isDecimal(unk interface{}) bool {
	x := unk.(string)
	_, err := strconv.ParseFloat(x, 64)
	if err != nil {
		return false
	}
	return true
}

func convertingDatatoString(v []interface{}) (string_type []string) {
	for _, data := range v {
		if !isNumeric(fmt.Sprintf("%s", data)) {
			string_type = append(string_type, fmt.Sprintf("%s", data))
		}
	}
	return
}

func convertingDatatoNumeric(v []interface{}) (numeric_type []int) {
	for _, data := range v {
		if isNumeric(fmt.Sprintf("%s", data)) {
			x := data.(string)
			val, err := strconv.Atoi(x)
			if err == nil {
				// fmt.Println("Berhasil di convert ke int dan masuk ke data")
				numeric_type = append(numeric_type, val)
			}
		}
	}
	//fmt.Println(numeric_type)
	return
}

func convertingDatatoDecimal(v []interface{}) (decimal_type []float64) {
	for _, data := range v {
		if isNumeric(fmt.Sprintf("%s", data)) && isDecimal(data.(interface{})) {
			x := data.(string)
			val, err := strconv.ParseFloat(x, 64)
			if err == nil {
				// fmt.Println("Berhasil di convert ke float dan masuk ke data")
				decimal_type = append(decimal_type, val)
			}
		}
	}
	//fmt.Println(decimal_type)
	return
}

func sortingColumn(sort_type int, column []interface{}) []bson.E {
	var sort_filter []bson.E
	for _, data := range column {
		sort_filter = append(sort_filter, bson.E{
			data.(string), sort_type,
		})
	}
	return sort_filter
}

func (repository *PilganRepository) GetAllDataPilgan(params map[string]interface{}) (*[]pilgan.PilganTone, error) {
	var limit int = 100
	var skip int = 0
	var filter []bson.E
	var sorting []bson.E
	for key, data := range params {
		if key != "limit" && key != "skip" && key != "sort_asc" && key != "sort_desc" {
			var query_in bson.E
			if isNumeric(fmt.Sprintf("%s", data.([]interface{})[0])) {
				query_in = bson.E{

					key, bson.D{
						{"$in", convertingDatatoDecimal(data.([]interface{}))},
					},
				}
			} else {
				query_in = bson.E{

					key, bson.D{
						{"$in", convertingDatatoString(data.([]interface{}))},
					},
				}
			}

			filter = append(filter, query_in)
		} else if key == "limit" {
			x := data.([]interface{})[0].(string)
			_, err := strconv.Atoi(x)
			if err == nil {
				limit, _ = strconv.Atoi(x)
			}
		} else if key == "skip" {
			x := data.([]interface{})[0].(string)
			_, err := strconv.Atoi(x)
			if err == nil {
				skip, _ = strconv.Atoi(x)
			}
		} else if key == "sort_asc" {
			sorting = append(sorting, sortingColumn(1, data.([]interface{}))...)
		} else if key == "sort_desc" {
			sorting = append(sorting, sortingColumn(-1, data.([]interface{}))...)
		}
	}
	projection := bson.M{"_id": 0}
	opts := options.Find().SetProjection(projection)
	opts.SetLimit(int64(limit))
	opts.SetSkip(int64(skip))

	if len(sorting) > 0 {
		opts.SetSort(sorting)
	}

	var data_filter interface{}
	if len(filter) > 0 {
		data_filter = filter
	} else {
		data_filter = bson.M{}
	}

	cursor, err := repository.client.Collection(repository.collection).Find(*repository.ctx, data_filter, opts)
	if err != nil {
		return nil, err
	}

	var list_pilgan []pilgan.PilganTone
	if err = cursor.All(*repository.ctx, &list_pilgan); err != nil {
		return nil, err
	}

	return &list_pilgan, nil
}
