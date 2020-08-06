# usm-ng-go 

### Its simple user management web ui, which is written in angular and  embedding as binary data into a go progra
 ####  * Steps to Create embedding as binary data 
  ####   * CMD: ng build 
  #### * build UI components to dist package 
  ####    * CMD: go-bindata-assetfs -o template.go dist/ ###
   #### package converts  into manageable Go source code. Useful for embedding binary data into a go program. The file data is optionally gzip compressed before being converted to a raw byte slice.

### Backend CRUD is written in Go gin webframework

#### Steps to build application
   ##### * go build
   ##### * usm-ng-go start

## Example
![alt text](https://github.com/kalaiarasan33/usm-ng-go/blob/master/usm.PNG)
