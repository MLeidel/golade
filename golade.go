package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var code1 string = `
package main

import (
    "os"
    //"fmt"

    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
)

// Create global vars for gtk widgets
`
var code2 string = `

func main() {
    // Create a new application. Change the appID string.
    application, err := gtk.ApplicationNew("com.change_this", glib.APPLICATION_FLAGS_NONE)
    errorCheck(err)

    // Connect Builder to application activate event
    application.Connect("activate", func() {

        // Get the GtkBuilder UI definition in the glade file.
        builder, err := gtk.BuilderNewFromFile("XXXXXX")
        errorCheck(err)

        // connect signals mapping handlers to callback functions
        signals := map[string]interface{} {`

var code3 string = `        }
        builder.ConnectSignals(signals)

`

var code4 string = `

        // show the window object with all widgets
        obj, err := builder.GetObject("window1")
        wnd := obj.(*gtk.Window)
        wnd.ShowAll()
        application.AddWindow(wnd)
    })

    // Launch the application
    os.Exit(application.Run(os.Args))
}


func errorCheck(e error) {
    if e != nil {
        // panic for any errors.
        panic(e)
    }
}

`

type wobjs struct {
	class string
	id    string
}

var objs []wobjs  // a slice of widget objects to obs
var sigs []string // a slice of signal handlers to item

func main() {

	if len(os.Args) < 2 {
        fmt.Println("Specify Glade file as argument")
		os.Exit(1)
	}

	gladefile := os.Args[1]

	f, err := os.Open(gladefile)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	o := wobjs{} // create variable to the wobjs struct
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, " id=") {
			pline := strings.Fields(line)
			class := strings.Split(pline[1], "\"")
			id := strings.Split(pline[2], "\"")
			o.class = class[1]
			o.id = id[1]
			objs = append(objs, o)
		}
		if strings.Contains(line, " handler=") {
			pline := strings.Fields(line)
			handl := strings.Split(pline[2], "\"")
			sigs = append(sigs, handl[1])
		}
	}
	check(scanner.Err())

    // Output begins here

	fmt.Println(code1)

	// Widget Globals
	for _, k := range objs {
		// k.class, k.id
		if !strings.Contains(k.class, "Button") {
			fmt.Println("var g_obj_" + k.id + " *gtk." + k.class[3:])
		}
	} // end for

    fmt.Println("\n\n//Signal Handler Functions \n")

	// Handler function bodies
	for _, k := range sigs {
		// handlers
		fmt.Println("func " + k + "() {\n\n}\n")
	}

    code2 = strings.Replace(code2,"XXXXXX", gladefile, 1)
	fmt.Println(code2)

	// Connect Signal Handlers
	for _, k := range sigs {
		// signals
		fmt.Println("            \"" + k + "\": " + k + ",")
	}

	fmt.Println(code3)

	// Global type assertions
	for _, k := range objs {
		if !strings.Contains(k.class, "Button") {
			fmt.Println("        obj_" + k.id + ", err := builder.GetObject(\"" + k.id + "\")")
			fmt.Println("        errorCheck(err)")
			fmt.Println("        g_obj_" + k.id + " = " + "obj_" + k.id + ".(*gtk." + k.class[3:] + ")")
			fmt.Println()
		}
	}

	fmt.Println(code4)
}

func check(e error) {
	if e != nil {
		// panic for any errors.
		panic(e)
	}
}
