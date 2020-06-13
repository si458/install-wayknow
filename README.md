# Custom Install of WaykNow

## Requirements
- [GoLang](https://golang.org/dl/)
- [Statik CLI](https://github.com/rakyll/statik) (embeds files)
- [rsrc CLI](https://github.com/akavel/rsrc) (embeds admin rights)
- [Latest WaykNow x64 MSI](https://wayk.devolutions.net/wayk-now/home/thankyou/wayksmsi64)
- [Custom Made WaykNow JSON Config File](https://helpwayk.devolutions.net/kb_configcommandline.html)

## Go Requirements
 - [configdir](https://github.com/kirsle/configdir) `go get github.com/kirsle/configdir`
 - [statik](https://github.com/rakyll/statik) `go get github.com/rakyll/statik`

## Instructions
 1. clone repo
 2. place the WaykNow X64 msi inside the `public` folder named as `wayknow.msi`
 3. edit the `WaykNow.cfg` which is located inside the `public` folder with your custom configs
 4. run command prompt or powershell at the clones repo folder
 5. run `statik` (this will build the embedded files required)
 6. run `rsrc -manifest install.exe.manifest -o install.syso` (this will embed the admin rights into the exe)
 7. **RUN STEP 8 OR 9, NOT BOTH!**
 8. run `go build -ldflags="-H windowsgui" -o install.exe` (this builds the app **WITHOUT A GUI**)
 9. run `go build -o install.exe` (this builds the app **WITH A BLACK GUI WINDOW FOR PROGRESS**)

