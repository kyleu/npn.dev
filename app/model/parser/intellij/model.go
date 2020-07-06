package parseintellij

type ijRoot struct {
	ID            int `xml:"id,attr"`
	ServerVersion string
	DateStyle     string
	StartupTime   int
}

func (ij *ijRoot) ParentID() int {
	return -1
}

type ijDatabase struct {
	ID        int    `xml:"id,attr"`
	Parent    int    `xml:"parent,attr"`
	Name      string `xml:"name,attr"`
	ObjectID  string `xml:"ObjectId"`
	Comment   string
	Owner     string
	Intro     int `xml:"IntrospectionStateNumber"`
	Current   int
	Relations string
}

func (ij *ijDatabase) ParentID() int {
	return ij.Parent
}

type ijSchema struct {
	ID          int    `xml:"id,attr"`
	Parent      int    `xml:"parent,attr"`
	Name        string `xml:"name,attr"`
	ObjectID    string `xml:"ObjectId"`
	Comment     string
	StateNumber int
	Owner       string
	Intro       int `xml:"IntrospectionStateNumber"`
	Current     int
}

func (ij *ijSchema) ParentID() int {
	return ij.Parent
}

type ijSequence struct {
	ID               int    `xml:"id,attr"`
	Parent           int    `xml:"parent,attr"`
	Name             string `xml:"name,attr"`
	ObjectID         string `xml:"ObjectId"`
	Owner            string
	StateNumber      int
	SequenceIdentity string
	CacheSize        int
	DataType         string
}

func (ij *ijSequence) ParentID() int {
	return ij.Parent
}

type ijObjectType struct {
	ID          int    `xml:"id,attr"`
	Parent      int    `xml:"parent,attr"`
	Name        string `xml:"name,attr"`
	ObjectID    string `xml:"ObjectId"`
	Owner       string
	StateNumber int
	SubKind     string
	Labels      string
	Definition  string
	SubCategory string
}

func (ij *ijObjectType) ParentID() int {
	return ij.Parent
}

type ijTable struct {
	ID          int    `xml:"id,attr"`
	Parent      int    `xml:"parent,attr"`
	Name        string `xml:"name,attr"`
	ObjectID    string `xml:"ObjectId"`
	Owner       string
	StateNumber int
}

func (ij *ijTable) ParentID() int {
	return ij.Parent
}

type ijColumn struct {
	ID                int    `xml:"id,attr"`
	Parent            int    `xml:"parent,attr"`
	Name              string `xml:"name,attr"`
	Position          int
	DataType          string
	NotNull           int
	StateNumber       int
	DefaultExpression string
	TypeID            int `xml:"TypeId"`
}

func (ij *ijColumn) ParentID() int {
	return ij.Parent
}

type ijIndex struct {
	ID          int    `xml:"id,attr"`
	Parent      int    `xml:"parent,attr"`
	Name        string `xml:"name,attr"`
	ObjectID    string `xml:"ObjectId"`
	StateNumber int
	ColNames    string
	Unique      int
	Primary     int
}

func (ij *ijIndex) ParentID() int {
	return ij.Parent
}

type ijKey struct {
	ID                  int    `xml:"id,attr"`
	Parent              int    `xml:"parent,attr"`
	Name                string `xml:"name,attr"`
	ObjectID            string `xml:"ObjectId"`
	StateNumber         int
	ColNames            string
	Primary             int
	UnderlyingIndexName string
}

func (ij *ijKey) ParentID() int {
	return ij.Parent
}

type ijForeignKey struct {
	ID              int    `xml:"id,attr"`
	Parent          int    `xml:"parent,attr"`
	Name            string `xml:"name,attr"`
	ObjectID        string `xml:"ObjectId"`
	StateNumber     int
	ColNames        string
	RefTableID      string `xml:"RefTableId"`
	RefColPositions string
	RefTableName    string
	RefKeyName      string
	RefColNames     string
}

func (ij *ijForeignKey) ParentID() int {
	return ij.Parent
}
