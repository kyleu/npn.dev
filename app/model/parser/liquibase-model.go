package parser

type lColumnConstraints struct {
	Nullable       bool   `xml:"nullable,attr"`
	PrimaryKey     bool   `xml:"primaryKey,attr"`
	PrimaryKeyName string `xml:"primaryKeyName,attr"`
}

type lColumn struct {
	Name        string             `xml:"name,attr"`
	T           string             `xml:"type,attr"`
	DefaultVal  string             `xml:"defaultValueComputed,attr"`
	Constraints lColumnConstraints `xml:"constraints"`
}

type lCreateTable struct {
	Name    string    `xml:"tableName,attr"`
	Columns []lColumn `xml:",any"`
}

type lAddForeignKeyConstraint struct {
	BaseColumnNames       string `xml:"baseColumnNames,attr"`
	BaseTableName         string `xml:"baseTableName,attr"`
	ConstraintName        string `xml:"constraintName,attr"`
	Deferrable            bool   `xml:"deferrable,attr"`
	InitiallyDeferred     bool   `xml:"initiallyDeferred,attr"`
	OnDelete              string `xml:"onDelete,attr"`
	OnUpdate              string `xml:"onUpdate,attr"`
	ReferencedColumnNames string `xml:"referencedColumnNames,attr"`
	ReferencedTableName   string `xml:"referencedTableName,attr"`
	Validate              bool   `xml:"validate,attr"`
}

type lAddUniqueConstraint struct {
	ColumnNames    string `xml:"columnNames,attr"`
	ConstraintName string `xml:"constraintName,attr"`
	TableName      string `xml:"tableName,attr"`
}

type lCreateIndex struct {
	Name      string    `xml:"indexName,attr"`
	TableName string    `xml:"tableName,attr"`
	Unique    bool      `xml:"unique,attr"`
	Columns   []lColumn `xml:",any"`
}
