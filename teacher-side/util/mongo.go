package util

import (
	"clientManagementSystem/config"
	"clientManagementSystem/module"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

var (
	StudentStatusCol *mongo.Collection
	StudentCol       *mongo.Collection
	HomeworkInfoCol  *mongo.Collection
)

const (
	STUDENTSTATUS = "studentStatus"
	STUDENTINFO   = "studentInfo"
	HOMEWORKINFO  = "homeworkInfo"
)

// connect to db and initialize the collection
func init() {

	client, err := mongo.NewClient(options.Client().
		ApplyURI(config.Config.DatabaseConfig.MongoConfig.DBAddress))
	if err != nil {
		log.Fatal(err)
	}

	expTime := time.Duration(config.Config.DatabaseConfig.MongoConfig.TimeOut) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), expTime)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	clientManagementSysDB := client.Database(config.Config.DatabaseConfig.MongoConfig.DBName)
	StudentStatusCol = clientManagementSysDB.
		Collection(config.Config.DatabaseConfig.MongoConfig.ClassCollection)
	StudentCol = clientManagementSysDB.
		Collection(config.Config.DatabaseConfig.MongoConfig.StudentCollection)
	HomeworkInfoCol = clientManagementSysDB.
		Collection(config.Config.DatabaseConfig.MongoConfig.HomeworkInfoCollection)

}

// trans values not ptr
func Save(colName string, content interface{}) (err error) {

	switch colName {
	case STUDENTSTATUS:
		{
			studentStatuses := content.([]module.StudentStatus)
			insertData := make([]interface{}, len(studentStatuses))
			for key, value := range studentStatuses {
				insertData[key] = value
			}
			_, err = StudentStatusCol.InsertMany(context.TODO(), insertData)
		}

	case STUDENTINFO:
		{
			studentInfo := content.([]module.StudentInfo)
			insertData := make([]interface{}, len(studentInfo))
			for key, value := range studentInfo {
				insertData[key] = value
			}
			_, err = StudentCol.InsertMany(context.TODO(), insertData)
		}

	case HOMEWORKINFO:
		{
			homeworkInfo := content.([]module.HomeworkInfo)
			insertData := make([]interface{}, len(homeworkInfo))
			for key, value := range homeworkInfo {
				insertData[key] = value
			}

			_, err = HomeworkInfoCol.InsertMany(context.TODO(), insertData)

		}

	}

	if err != nil {
		return err
	}

	return nil
}

func UpdateOne(colName string, content interface{}) (err error) {

	//opt := options.Update().SetUpsert(true)

	switch colName {
	case STUDENTSTATUS:
		{
			studentStatus := content.(module.StudentStatus)

			filter := bson.M{
				"studentinfo.student_id": studentStatus.StudentId,
				"class.class_name":       studentStatus.ClassName,
				"class.class_start_date": studentStatus.ClassStartDate,
			}
			update := bson.M{"$set": studentStatus}
			_, err = StudentStatusCol.UpdateOne(context.TODO(), filter, update)

		}
		//
		//case STUDENTINFO:{
		//	studentInfo := content.(StudentInfo)
		//	filter :=
		//}
	}

	if err != nil {
		return err
	}

	return nil
}

// find all filter must contain the studentId, ClassName, ClassStartDate
// match student password should just contain studentId
// the return interface should be transform into StudentStatus or StudentInfo
//
// for example:
//
// filterMap:= map[string]string{
//		"StudentId" = "Uxx",
//		"ClassName" = "shixun",
//		"ClassStartDate" = "555555",
// }
//
// in Case of homeworkInfo, the filterMap should contains such a key: "HomeworkName"
func FindOne(colName string, filterMap map[string]string) (result interface{}, err error) {

	switch colName {
	case STUDENTSTATUS:
		{
			var (
				classStartDate int64
				results        module.StudentStatus
			)
			classStartDate, err = strconv.ParseInt(filterMap["ClassStartDate"], 10, 64)
			filter := bson.M{
				"studentinfo.student_id": filterMap["StudentId"],
				"class.class_name":       filterMap["ClassName"],
				"class.class_start_date": classStartDate,
			}

			err = StudentStatusCol.FindOne(context.TODO(), filter).Decode(&results)
			if err != nil {
				return nil, err
			}

			return results, nil

		}

	case STUDENTINFO:
		{
			var results module.StudentInfo
			filter := bson.M{
				"student_id": filterMap["StudentId"],
			}

			err := StudentCol.FindOne(context.TODO(), filter).Decode(&results)
			if err != nil {
				return nil, err
			}

			return results, nil
		}

	case HOMEWORKINFO:
		{
			var result module.HomeworkInfo
			filter := bson.M{
				"homework_name": filterMap["HomeworkName"],
			}

			err := HomeworkInfoCol.FindOne(context.TODO(), filter).Decode(result)
			if err != nil{
				return nil, err
			}

			return result, nil
		}

	}

	return result, nil
}

// colValue is just used for search the specific person's all tips, so the value is StudentId
// the return interface should be just transform into [] StudentStatus or [] StudentInfo
func FindAll(colName string, studentId string, className ...string) (result interface{}, err error) {

	switch colName {
	case STUDENTSTATUS:
		{

			cursor, err := StudentStatusCol.Find(context.TODO(),
				bson.M{"studentinfo.student_id": studentId, "class.class_name": className[0]})
			if err != nil {
				return nil, err
			}

			var results []module.StudentStatus
			if err = cursor.All(context.TODO(), &results); err != nil {
				return nil, err
			}

			return results, nil
		}

	case STUDENTINFO:

		cursor, err := StudentCol.Find(context.TODO(), bson.M{})
		if err != nil {
			return nil, err
		}

		var results []module.StudentInfo
		if err = cursor.All(context.TODO(), &results); err != nil {
			return nil, err
		}

		return results, nil

	case HOMEWORKINFO:

		cursor, err := HomeworkInfoCol.Find(context.TODO(), bson.M{})
		if err != nil{
			return nil, err
		}

		var results []module.HomeworkInfo
		if err = cursor.All(context.TODO(), &result); err != nil{
			return nil, err
		}

		return results, nil
	}


	return nil, nil
}
