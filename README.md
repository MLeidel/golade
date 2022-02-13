# golade
*generates Go source code for a gtk GUI from a Glade xml file*

`golade` creates a Go program using your Glade (xml) definition file.  
It sets up the global widget objects, signal connections and  
handler function bodies.  

Usage: `golade my.glade > my.go`  

The UI should run (display) at this point; without fully functioning widgets.

Run: `go run my.go`


