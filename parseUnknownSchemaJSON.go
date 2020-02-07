package main
import (

"encoding/json"
"fmt"
)

var jsonBS = []byte(`{
   "mykey": "myValue",
  "data": [{
    "type": "articles",
    "id": "1",
    "attributes": {
      "title": "JSON:API paints my bikeshed!",
      "body": "The shortest article. Ever.",
      "created": "2015-05-22T14:56:29.000Z",
      "updated": "2015-05-22T14:56:28.000Z"
    },
    "relationships": {
      "author": {
        "data": {"id": "42", "type": "people"}
      }
    }
  }],
  "included": [
    {
      "type": "people",
      "id": "42",
      "attributes": {
        "name": "John",
        "age": 80,
        "gender": "male"
      }
    }
  ]
}`)

func main(){
	//declare an interface that u can use to unmarshal the JSON into
	var genIface interface{}
	err := json.Unmarshal(jsonBS, &genIface)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(genIface)
	// We can access next  level data by using a map[string]interface. All top level keys can be accessed in this manner
	myData := genIface.(map[string]interface{})
	fmt.Println(myData["mykey"])
	printJSON(genIface)
}

func printJSON(i interface{}){
	switch val := i.(type){
	case string:
		fmt.Println("This is a string -- ", val)
	case float64:
		fmt.Println("is a float64 -- ", val)
	case []interface{}:
		fmt.Print("is an array  -- ", val)
		for ind, value  := range val{
			fmt.Println("This is an array. index -- ", ind)
			printJSON(value)
		}
	case map[string]interface{}:
		fmt.Print("this is a map object:")
		for key, value  := range val {
			fmt.Println("This is an map key -- ", key )
			printJSON(value)
		}
	default:
		fmt.Println("unknown type -- ", val)
	}

}
