# appConfig 

appConfig is a json config file initializer 

## Usage 

```json
{
        "field1" : "data",
        "field2" : [1,2],
        "field3" : 1234678
}

```

```go 
import "github.com/PierreKieffer/appConfig"

type Configuration struct {
        Field1  string `json:"field1"`
        Field2 []int `json:"field2"`
        Field3 int64 `json:"field3"`
}

func main() {
        configuration := Configuration{}
        appConfig.InitConfig("config.json", &configuration)
}

```
