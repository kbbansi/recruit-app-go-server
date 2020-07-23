package util

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"time"
)

const (
	// numberRegex is the regular expression for obtaining numbers only in a string.
	numberRegex = "[0-9]+"

	// dateLayout is the date layout for ISO 8601 date format (YYYY-MM-DD) used for parsing date sent in a string.
	dateLayout = "2006-01-02"

	// timeRegex the regular expression for the ISO 8601 extended time format (hh:mm:ss) used for parsing time sent in a string.
	timeRegex = `([01]\d|2[0-3]):([0-5]\d):([0-5]\d)`

	// dateTimeLayout is the date time layout for the mysql/mariadb datetime type.
	dateTimeLayout = "2006-01-02 15:04:05"

	// letterBytes is the charset used for generating random client ids for subscribers.
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	// letterIdxBits is the number of bits needed to represent the index of a character from letterBytes in binary
	letterIdxBits = 6

	// letterIdxMask is the binary representation of 63. ie `111111`. This is used as a mask to get only the last
	// 6 bits of a int63 number
	letterIdxMask = 1<<letterIdxBits - 1

	// AMSExchangeName is the name of the exchange used for the project.
	AMSExchangeName = "AMS_EXCHANGE"

	// BatchFileQueueName is the name of the queue for file uploads.
	BatchFileQueueName = "BATCH_FILE_QUEUE"

	// BatchFileQueueRoutingKey is the routing key for the file upload queue.
	BatchFileQueueRoutingKey = "file"

	// SchemaGeneratorQueueRoutingKey is the routing key for the schema generator queue.
	SchemaGeneratorQueueRoutingKey = "schema"

	// SchemaGeneratorQueueName is the name of the queue for the schema generation for organizations.
	SchemaGeneratorQueueName = "SCHEMA_GENERATOR_QUEUE"

	// EmployeeBatchCreationEventType is the event type for an employee batch upload
	EmployeeBatchCreationEventType = "EMPLOYEE_BATCH_CREATION"

	// ClockInformationBatchCreationEventType is the event type for clock information upload
	ClockInformationBatchCreationEventType = "CLOCK-INFORMATION_BATCH_CREATION"

	// OrganizationTable is the table for the organization information of an organization.
	OrganizationTable = "organization_table"

	// EmployeeTable is the table for the employee information of an organization.
	EmployeeTable = "employee_table"

	// DepartmentTable is the table for the department information of an organization.
	DepartmentTable = "department_table"

	// BranchTable is the table for the branch information of an organization.
	BranchTable = "branch_table"

	// ClockLogTable is the table for the clock-log information of an organization.
	ClockLogTable = "clock_log_table"

	// ConfigTable is the table for the config information of an organization.
	ConfigTable = "config_table"

	// HolidayTable is the table for the holiday information of an organization.
	HolidayTable = "holiday_table"

	// OrganizationSchemaFile is the file name for the definition of the organization table.
	OrganizationSchemaFile = "organization.sql"

	// EmployeeScheamFile is the file name for the definition of the employee table.
	EmployeeSchemaFile = "employee.sql"

	// DepartmentSchemaFile is the file name for the definition of the department table.
	DepartmentSchemaFile = "department.sql"

	// BranchSchemaFile is the file name for the definition of the branch table.
	BranchSchemaFile = "branch.sql"

	// ClockLogSchemaFile is the file name for the definition of the clock log table.
	ClockLogSchemaFile = "clock_log.sql"

	// ConfigSchemaFile is the file name for the definition of the config table.
	ConfigSchemaFile = "config.sql"

	// HolidaySchemaFile is the file name for the definition of the holiday table.
	HolidaySchemaFile = "holiday.sql"

	// OrganiztionTablePlaceholder is the name that is appended to an organization's namespace to generate an organization's organization table name.
	OrganizationTablePlaceholder = "_organization"

	// EmployeeTablePlaceholder is the name that is appended to an organization's namespace to generate an organization's employee table name.
	EmployeeTablePlaceholder = "_employee"

	// DepartmentTablePlaceholder is the name that is appended to an organization's namespace to generate an organization's department table name.
	DepartmentTablePlaceholder = "_department"

	// BranchTablePlaceholder is the name that is appended to an organization's namespace to generate an organization's branch table name.
	BranchTablePlaceholder = "_branch"

	// ClockLogTablePlaceholder is the name that is appended to an organization's namespace to generate an organization's clock-log table name.
	ClockLogTablePlaceholder = "_clock_log"

	// ConfigTablePlaceholder is the name that is appended to an organization's namespace to generate an organization's config table name.
	ConfigTablePlaceholder = "_config"

	// HolidayTablePlaceholder is the name that is appended to an organization's namespace to generate an organization's holiday table name.
	HolidayTablePlaceholder = "_holiday"

	//API_ERRORS is the name of error file for api errors
	API_ERRORS = "api-errors.txt"

	//NOTIFICATIONS_ERRORS is the name of the error file for notification errors of the Ghana Post Management API
	NOTIFICATIONS_ERRORS = "notifications-errors.txt"

	//LOGS is the name of the file for keeping logs
	LOGS = "logs.txt"
)

var (
	// Encoder is the encoder used by the tests to encode json request data.
	Encoder = new(bytes.Buffer)

	// ID is used by the tests to enable the use of resource ids over different tests.
	ID int

	// Count is used by the handlers to count instances of
	Count int

	// used as middle ground between datetime and UTC.
	middleman time.Time

	// CreatedOn is used as a holder for the created on time of a resource before it is reformatted in UTC.
	CreatedOn string

	// ModifiedOn is used as a holder for the modified on time of a resource before it is reformatted in UTC.
	ModifiedOn string

	// NameSpace holds the name space of an organization.
	NameSpace string

	//Password is the user defined password of an account
	Password string
)

// Data is the struct for the data in pdSuccess.
type Data struct {
	ID         int    `json:"id"`
	ActionType string `json:"actionType"`
}

// PDSuccess is the response struct for successful DELETE, POST and PUT methods.
type PDSuccess struct {
	Status int `json:"status"`
	Data   *Data  `json:"data"`
}

// Fail is the response struct for all failures.
type Fail struct {
	Status int `json:"status"`
	Reason string `json:"reason"`
}

//Found is the response for informational success
type MultiPurpose struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	User     string `json:"user, omitempty"`
	UserID   int    `json:"userID, omitempty"`
	ToUserID int    `json:"toUserID, omitempty"`
	ToUser   string `json:"toUser, omitempty"`
}

// APILogMessage is the payload stuct for an api log message.
type APILogMessage struct {
	Method      string `json:"method"`
	URI         string `json:"uri"`
	Name        string `json:"name"`
	TimeElapsed int64  `json:"time_elapsed"`
	Service     string `json:"service"`
	Type        string `json:"type"`
}

// APIErrorMessage is the payload struct for an api error message.
type APIErrorMessage struct {
	Message string `json:"message"`
	Service string `json:"service"`
	Type    string `json:"type"`
}

func ConvertDateTimeStringToUTCString(dateTime string) (string, error) {
	// incase datetime value is NULL
	if dateTime == "" {
		return dateTime, nil
	}
	middleman, err := time.Parse(dateTimeLayout, dateTime)
	if err != nil {
		return "", err
	}

	return middleman.Format(time.RFC3339), nil
}

func ConvertStringToInt(value string) (int, error) {
	number, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0, err
	}
	fmt.Println(number)
	return int(number), nil
}

//HashPassword hashes a password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Printf(err.Error())
	}
	return err == nil
}
