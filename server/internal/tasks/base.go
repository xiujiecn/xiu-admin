package tasks

var (
	innerTaskNames = map[string]string{}
)

func RegisterInnerTask(name string, methodName string) {
	innerTaskNames[name] = methodName
}

func GetInnerTaskNames() map[string]string {
	return innerTaskNames
}

func GetInnerTaskName(name string) string {
	return innerTaskNames[name]
}
