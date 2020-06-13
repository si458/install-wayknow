package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"

	_ "./statik"
	"github.com/kirsle/configdir"
	"github.com/rakyll/statik/fs"
)

// go build -ldflags -H=windowsgui install.go

func logStuff(mylog string) {
	if true {
		log.Println(mylog)
	}
}

func getAndWrite(statikFS http.FileSystem, getname string, writename string) {
	var thefile []byte
	var err error
	logStuff("Getting Packaged File")
	if thefile, err = fs.ReadFile(statikFS, getname); err != nil {
		log.Fatal(err)
	}
	logStuff("Writing Packaged File")
	if err = ioutil.WriteFile(writename, thefile, 0777); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var statikFS http.FileSystem
	var err error
	var mycode int
	var cmd *exec.Cmd
	logStuff("Getting Wayk Global Folder")
	globalPath := strings.Join(configdir.SystemConfig("Wayk"), "")
	logStuff("Getting Wayk Roaming Folder")
	localRoaming := configdir.LocalConfig("Wayk")
	logStuff("Getting Temp Folder")
	tempPath := os.TempDir()
	logStuff("Loading Packaged Files")
	if statikFS, err = fs.New(); err != nil {
		log.Fatal(err)
	}
	getAndWrite(statikFS, "/wayknow.msi", tempPath+"\\wayknow.msi")
	logStuff("Running Install MSI Command")
	abc := fmt.Sprintf(` /i %v /passive`, tempPath+"\\wayknow.msi") // Leave a space at the beginning
	//abc := fmt.Sprintf(` /i %v /passive INSTALLDESKTOPSHORTCUT="" INSTALLSTARTMENUSHORTCUT=""`, tempwayknow) // Leave a space at the beginning
	msi := exec.Command("msiexec")
	msi.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    false,
		CmdLine:       abc,
		CreationFlags: 0,
	}
	if err = msi.Run(); err != nil {
		log.Fatal("Error: ", err)
	}
	logStuff("Removing Temp Wayk MSI")
	if err = os.RemoveAll(tempPath + "\\wayknow.msi"); err != nil {
		log.Fatal(err)
	}
	logStuff("Stop Wayk Service")
	cmd = exec.Command("net", "stop", "WaykNowService")
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
	if err = cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			mycode = exitError.ExitCode()
		}
		if mycode != 2 {
			log.Fatal(err)
		}
	}
	logStuff("Checking Wayk Global Folder")
	if _, err = os.Stat(globalPath); os.IsNotExist(err) {
		logStuff("Wayk Global Folder Doesnt Exist")
		if err = os.MkdirAll(globalPath, 0777); err != nil {
			log.Fatal(err)
		}
		logStuff("Created Wayk Global Folder")
	} else {
		logStuff("Wayk Global Folder Already Exists")
	}
	logStuff("Remove Existing Local Wayk Install")
	if err = os.RemoveAll(localRoaming); err != nil {
		log.Fatal(err)
	}
	getAndWrite(statikFS, "/WaykNow.cfg", globalPath+"\\WaykNow.cfg")
	logStuff("Start Wayk Service")
	cmd = exec.Command("net", "start", "WaykNowService")
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
	if err = cmd.Run(); err != nil {
		log.Fatal(err)
	}
	logStuff("Start Application")
	cmd = exec.Command(`C:\Program Files\Devolutions\Wayk Now\WaykNow.exe`)
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
