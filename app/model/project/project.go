package project

type Project struct {
	Key        string
	SchemaKeys []string
}

var TestProject = Project{
	Key:        "test",
	SchemaKeys: nil,
}
