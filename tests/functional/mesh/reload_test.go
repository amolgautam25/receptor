package mesh

import (
	"fmt"
	"testing"

	_ "github.com/fortytw2/leaktest"
	"github.com/project-receptor/receptor/tests/functional/lib/mesh"
)

func TestReload(t *testing.T) {
	t.Parallel()

	// declare the data structure that will hold the node Yaml data
	data := mesh.YamlData{}
	data.Nodes = make(map[string]*mesh.YamlNode)

	// Seed the datastructure with 3 nodes
	data.Nodes["node2"] = &mesh.YamlNode{
		Connections: map[string]mesh.YamlConnection{},
		Nodedef: []interface{}{
			map[interface{}]interface{}{
				"tcp-listener": map[interface{}]interface{}{
					"cost": 1.0,
					"nodecost": map[interface{}]interface{}{
						"node1": 1.0,
						"node3": 1.0,
					},
				},
			},
		},
	}
	data.Nodes["node1"] = &mesh.YamlNode{
		Connections: map[string]mesh.YamlConnection{},
		Nodedef:     []interface{}{},
	}
	data.Nodes["node3"] = &mesh.YamlNode{
		Connections: map[string]mesh.YamlConnection{},
		Nodedef:     []interface{}{},
	}

	// start the mesh with the above data.

	m, err := mesh.NewCLIMeshFromYaml(data, "reload_mesh_test")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(m)
}
