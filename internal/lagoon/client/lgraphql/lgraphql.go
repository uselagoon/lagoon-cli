// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package lgraphql generated by go-bindata.// sources:
// _lgraphql/addBillingGroup.graphql
// _lgraphql/addEnvVariable.graphql
// _lgraphql/addGroup.graphql
// _lgraphql/addGroupsToProject.graphql
// _lgraphql/addNotificationEmail.graphql
// _lgraphql/addNotificationMicrosoftTeams.graphql
// _lgraphql/addNotificationRocketChat.graphql
// _lgraphql/addNotificationSlack.graphql
// _lgraphql/addNotificationToProject.graphql
// _lgraphql/addOrUpdateEnvironment.graphql
// _lgraphql/addProject.graphql
// _lgraphql/addProjectToBillingGroup.graphql
// _lgraphql/addSshKey.graphql
// _lgraphql/addUser.graphql
// _lgraphql/addUserToGroup.graphql
// _lgraphql/environmentByName.graphql
// _lgraphql/me.graphql
// _lgraphql/projectByName.graphql
package lgraphql

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __lgraphqlAddbillinggroupGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x89\x24\x97\x16\x15\xa5\xe6\x25\x57\x5a\x29\x38\x43\x59\x10\xf1\xa4\xcc\x9c\x9c\xcc\xbc\xf4\xe0\xfc\xb4\x92\xf2\xc4\x22\xb8\x26\x4d\x85\x6a\x2e\x05\x05\x05\x85\xc4\x94\x14\x27\x88\x0a\xf7\xa2\xfc\xd2\x02\x8d\xcc\xbc\x82\xd2\x12\x2b\xa8\xa4\x82\x02\xc4\x1e\xb0\x75\x3a\x50\x21\x84\x45\x70\x3b\x61\x52\x18\x76\xa1\xdb\x0e\x56\x57\xab\x09\x37\x3e\x33\x05\xc9\x1e\x88\x24\x17\x08\x03\x02\x00\x00\xff\xff\x90\xae\x83\xc3\xed\x00\x00\x00")

func _lgraphqlAddbillinggroupGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddbillinggroupGraphql,
		"_lgraphql/addBillingGroup.graphql",
	)
}

func _lgraphqlAddbillinggroupGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddbillinggroupGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addBillingGroup.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddenvvariableGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\xb1\xca\xc3\x20\x14\x46\xf7\x3c\xc5\xf7\xc3\x3f\x24\x90\x27\x70\xef\x90\x39\xa5\xfb\x6d\x95\x22\x24\x37\x21\xbd\x06\x42\xe9\xbb\x17\xbd\xda\x4a\x07\x97\x73\xd4\xef\xcc\x41\x48\xfc\xc2\x68\x1b\xe0\x5f\x8e\xd5\x19\x9c\x78\xbf\xd0\xe6\xe9\x3a\xb9\xf3\xb1\xba\xbf\xbe\xa8\xc1\x1a\x0c\x2c\x0a\x1e\xb7\xe5\xe7\xf2\x18\x89\x4a\xa6\xd9\x19\x8c\xb2\x79\xbe\x2b\xd9\x69\x0a\x5f\xd4\xe1\xd9\x00\x00\x59\x5b\x7d\xd0\x7a\x5e\x83\x98\xec\x00\xcd\x49\xd3\x7d\x85\x62\x46\xee\x29\x38\xc7\x68\x54\x81\x1a\x91\x5a\x0a\xca\x15\x5a\x93\xd8\xab\xfb\xcc\x79\x5b\x3d\x54\xd9\xc4\xf3\x0e\x00\x00\xff\xff\x1d\x4f\x13\xd7\x24\x01\x00\x00")

func _lgraphqlAddenvvariableGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddenvvariableGraphql,
		"_lgraphql/addEnvVariable.graphql",
	)
}

func _lgraphqlAddenvvariableGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddenvvariableGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addEnvVariable.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddgroupGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x4c\x49\x71\x2f\xca\x2f\x2d\xd0\xc8\xcc\x2b\x28\x2d\xb1\x82\x8a\x2a\x28\x40\x14\x83\xf5\x80\x45\x6a\x35\xe1\x52\x99\x29\x48\x6a\x20\x92\x5c\x20\x0c\x08\x00\x00\xff\xff\x33\xcc\xe8\xad\x6e\x00\x00\x00")

func _lgraphqlAddgroupGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddgroupGraphql,
		"_lgraphql/addGroup.graphql",
	)
}

func _lgraphqlAddgroupGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddgroupGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addGroup.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddgroupstoprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x29\x28\xca\xcf\x4a\x4d\x2e\xb1\x52\x08\x80\x30\x3c\xf3\x0a\x4a\x4b\x14\x75\x40\x52\xe9\x45\xf9\xa5\x05\xc5\x56\x0a\xd1\xee\x20\x06\x44\x22\x56\x51\x53\xa1\x9a\x4b\x41\x41\x41\x21\x31\x25\x05\x2c\x5e\x1c\x92\x0f\xd5\xaa\x91\x09\x52\x62\x05\x95\x57\x50\x80\x1b\x0d\xb3\x04\x2a\x0e\x33\x17\x6a\x01\x58\xb4\x56\x13\xae\x2d\x33\x05\xca\xc8\x4b\xcc\x4d\x85\x48\x72\x81\x30\x20\x00\x00\xff\xff\x8b\x70\xd6\xd3\xb8\x00\x00\x00")

func _lgraphqlAddgroupstoprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddgroupstoprojectGraphql,
		"_lgraphql/addGroupsToProject.graphql",
	)
}

func _lgraphqlAddgroupstoprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddgroupstoprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addGroupsToProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddnotificationemailGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x89\xa4\xe6\x26\x66\xe6\x38\xa6\xa4\x14\xa5\x16\x17\xc3\x65\x34\x15\xaa\xb9\x14\x14\x14\x14\x12\x53\x52\xfc\xf2\x4b\x32\xd3\x32\x93\xc1\x66\xb8\x82\xd4\x6a\x64\xe6\x15\x94\x96\x58\x41\x55\x28\x28\x40\x8c\x04\x9b\xac\x03\x15\x42\x35\x13\xc5\x0a\xb0\x8a\x5a\x4d\xb8\xee\xcc\x14\x24\x63\x20\x92\x5c\x20\x0c\x08\x00\x00\xff\xff\x7d\x3b\x64\xad\xb7\x00\x00\x00")

func _lgraphqlAddnotificationemailGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddnotificationemailGraphql,
		"_lgraphql/addNotificationEmail.graphql",
	)
}

func _lgraphqlAddnotificationemailGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddnotificationemailGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addNotificationEmail.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddnotificationmicrosoftteamsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x89\x94\xa7\x26\x65\xe4\xe7\x67\xc3\x05\x35\x15\xaa\xb9\x14\x14\x14\x14\x12\x53\x52\xfc\xf2\x4b\x32\xd3\x32\x93\xc1\xda\x7d\x33\x93\x8b\xf2\x8b\xf3\xd3\x4a\x42\x52\x13\x73\x8b\x35\x32\xf3\x0a\x4a\x4b\xac\xa0\x4a\x15\x14\x20\xc6\x82\x4d\xd7\x81\x0a\xc1\xcd\x85\xd9\x00\x16\xaf\xd5\x84\xeb\xc9\x4c\x41\xd2\x0c\x91\xe4\x02\x61\x40\x00\x00\x00\xff\xff\x97\x56\x46\x93\xb1\x00\x00\x00")

func _lgraphqlAddnotificationmicrosoftteamsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddnotificationmicrosoftteamsGraphql,
		"_lgraphql/addNotificationMicrosoftTeams.graphql",
	)
}

func _lgraphqlAddnotificationmicrosoftteamsGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddnotificationmicrosoftteamsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addNotificationMicrosoftTeams.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddnotificationrocketchatGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8d\x31\x0e\xc2\x30\x0c\x45\xf7\x9c\xe2\x23\x31\xb4\x52\x4f\xd0\x95\x9d\x01\x4e\x60\x92\x40\xac\x52\x1b\x21\x57\x0c\x88\xbb\xa3\xa6\x49\x04\x83\x97\xf7\xad\xf7\xe6\xc5\xc8\x58\x05\x9d\x03\xf6\x42\x73\x1c\x71\xb6\x27\xcb\x6d\x37\xac\xc4\x27\x12\x89\xf7\x7f\xf8\x8a\x97\xa4\x3a\x35\xd8\xe3\xed\x00\x80\x42\x38\xaa\xf1\x95\x7d\x76\x9e\xd4\x4f\xd1\x0e\x89\xac\x63\x79\x2c\x36\x96\x37\x60\xeb\xe4\xdc\x50\x50\x0b\xd5\x64\x1d\x5a\xac\x66\x33\xff\xf4\x4d\xc6\xe1\xc7\xba\x8d\x6e\xbd\x6f\x00\x00\x00\xff\xff\x53\x0f\xa2\x30\xdb\x00\x00\x00")

func _lgraphqlAddnotificationrocketchatGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddnotificationrocketchatGraphql,
		"_lgraphql/addNotificationRocketChat.graphql",
	)
}

func _lgraphqlAddnotificationrocketchatGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddnotificationrocketchatGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addNotificationRocketChat.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddnotificationslackGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8d\x41\x0e\x82\x40\x0c\x45\xf7\x73\x8a\x6f\xe2\x02\x12\x4e\xc0\x21\xdc\x70\x82\x3a\x33\x4a\x03\xb4\xc6\x94\xb8\x30\xde\xdd\x30\x30\x8d\x2e\xba\x79\xbf\x79\x6f\x59\x8d\x8c\x55\xd0\x04\xe0\x2c\xb4\xe4\x1e\x83\x3d\x59\xee\xa7\x6e\x23\x71\x24\x91\x3c\xff\xc3\x57\xbe\x8e\xaa\x93\xc3\x16\xef\x00\x00\x94\xd2\x45\x8d\x6f\x1c\x8b\x73\x98\x29\x4e\x0d\xcb\x63\xb5\xfe\xf8\x00\xf6\x44\x29\x75\x07\xf2\x46\xad\xd5\xc1\x3b\xb5\x58\xf8\xa7\x75\x19\xa7\x1f\xeb\x3e\x86\xed\xbe\x01\x00\x00\xff\xff\xd9\x6e\xb4\xda\xd6\x00\x00\x00")

func _lgraphqlAddnotificationslackGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddnotificationslackGraphql,
		"_lgraphql/addNotificationSlack.graphql",
	)
}

func _lgraphqlAddnotificationslackGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddnotificationslackGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addNotificationSlack.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddnotificationtoprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x29\x28\xca\xcf\x4a\x4d\x2e\xb1\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x09\xe6\xe5\x97\x64\xa6\x65\x26\x83\x95\x85\x54\x16\xa4\x5a\x29\xf8\xa1\x89\x60\xaa\xf3\x4b\xcc\x4d\x85\x9b\xa2\xa9\x50\xcd\xa5\xa0\xa0\xa0\x90\x98\x92\x82\xa2\x33\x3f\x00\x62\x9d\x46\x66\x5e\x41\x69\x89\x15\x54\x95\x82\x02\xdc\x15\x30\xf7\x40\xc5\x31\x1d\x82\xe1\x36\x2c\x2a\x21\x4e\xc1\x70\x1d\x58\x65\xad\x26\xdc\xd2\xcc\x14\x98\x5e\xb8\x24\x17\x08\x03\x02\x00\x00\xff\xff\x5d\xba\xcd\xcd\x21\x01\x00\x00")

func _lgraphqlAddnotificationtoprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddnotificationtoprojectGraphql,
		"_lgraphql/addNotificationToProject.graphql",
	)
}

func _lgraphqlAddnotificationtoprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddnotificationtoprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addNotificationToProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddorupdateenvironmentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xcd\x4a\xc5\x30\x10\x85\xf7\x7d\x8a\x11\x5c\xdc\x0b\xf7\x09\xb2\x14\x0b\xba\x51\xf1\xe7\x01\x82\x99\x6a\xa4\x9d\x84\x74\x5a\x28\xe2\xbb\x4b\x9b\x26\x9d\xf4\x76\x51\x28\xe7\x9c\xcc\xcc\xf9\xba\x81\x35\x5b\x47\x70\xaa\x00\x6e\x49\x77\xa8\xe0\x8d\x83\xa5\xaf\x9b\xcb\xac\xf8\xe0\x7e\xf0\x93\x15\x3c\x12\x47\xc5\xa0\x6f\xdd\xf4\x3e\x79\x54\x70\x9f\xff\xa5\x77\xa7\x7b\x7c\xc5\xa6\x1c\x14\xad\x07\xd4\x46\x58\x72\xa0\xe5\x16\x0b\x1d\x69\xb4\xc1\x51\x87\xc4\x71\x5b\x4d\xe3\xb6\xca\x79\xa4\xfe\xdb\x36\xfc\x12\x2f\x7c\x92\xa7\x9f\xe1\xb7\x02\x00\xd0\xc6\x3c\x87\x0f\x6f\x34\x63\xbd\x4d\x3b\x59\xf2\x03\xab\x35\x03\x10\x5b\x2f\xe5\x2f\xab\x94\x6b\x27\x00\xc9\x90\xed\x05\x8a\xd2\xce\x00\x4a\x20\x65\x28\xa3\x28\xd1\xec\x16\x45\x2a\x92\x51\x0a\x5c\xe1\xd9\x03\x4b\xc1\x63\x52\x87\x00\x97\x17\x7f\xe7\x4c\xc6\x1a\x81\x28\x9a\xd5\xfc\xfd\x07\x00\x00\xff\xff\x2f\xb0\x49\x31\x36\x02\x00\x00")

func _lgraphqlAddorupdateenvironmentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddorupdateenvironmentGraphql,
		"_lgraphql/addOrUpdateEnvironment.graphql",
	)
}

func _lgraphqlAddorupdateenvironmentGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddorupdateenvironmentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addOrUpdateEnvironment.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\x4b\x6e\x2a\x31\x10\x45\xe7\xac\xc2\x48\x6f\x00\x12\x2b\x60\xfa\x92\x01\x4a\x06\x28\x24\x0b\x30\xed\x02\x9c\xd8\xae\x8e\x5d\x6e\x09\x45\xd9\x7b\xd4\x1f\xff\xc0\xf4\xa0\x27\xf7\x56\xb5\xac\x3a\x47\x7b\xe2\x24\xd1\xb0\xd5\x82\xb1\x7f\x86\x6b\xd8\xb2\x03\x59\x69\xce\xcb\x4d\x9f\x9c\x25\x7d\x58\x55\x66\xce\x1f\x4f\xa8\x04\xd8\x10\x0f\x29\xb6\x60\xdc\x45\x9e\x68\xcb\x76\x86\x96\x65\xb6\xb7\xf8\x09\x0d\xed\x39\x11\x58\x53\xec\xf1\x86\x64\x07\x87\xab\x23\xd0\xee\x09\x5a\x85\xd7\xc7\xfd\xde\xa2\x46\x82\xc7\x03\x6f\xa0\xb1\x9b\xe9\xdf\xb9\xfb\x2a\xda\xa3\xe5\xa6\xb9\x80\x2b\xc2\xd6\x2b\x65\xe1\xdb\x83\xa3\x9b\xc2\xa2\xf0\x4d\x7f\xb0\x67\xd3\x49\x8b\x46\x83\xa1\xf2\x3a\xdc\x13\xee\x84\x82\xe1\x0c\xe3\xbd\x08\x2d\x3f\xc3\x7f\xae\x9a\x14\x0a\xe8\x40\x61\xdb\xef\x67\xbf\x72\xaf\x52\x4b\x4a\x53\xad\x95\x1d\x27\x78\x81\x78\x93\x35\xfb\x59\x30\xc6\x18\x17\x62\x3a\xea\x4a\x9a\xd6\xd3\x76\xca\x19\x1b\x21\x0e\x2c\x37\x53\x14\x28\x4e\x38\x43\x9c\x81\x4c\x50\x43\x99\xf1\x4c\x1c\xef\xca\x5b\xb0\x8f\x90\x87\xc5\x2a\xed\x9a\x03\xd5\x85\x88\xbf\x6a\x45\x75\x25\x08\x51\xd3\xa4\xba\x30\x1a\x72\x6f\x4d\x18\x4e\xc2\x44\x77\x42\x55\x6a\x53\x58\x14\x47\xea\x02\xd5\xc5\x8a\xef\x8b\x4a\x45\xbb\x22\xc2\xdc\xad\xdc\xb4\x30\x30\xef\xd9\xac\x86\xe9\xcd\x49\xc2\xcc\xc8\xa1\xfd\x5d\x47\xef\xa4\xc8\x04\x1c\xcb\x45\xff\xfd\x05\x00\x00\xff\xff\xd3\x71\xbe\x59\x63\x04\x00\x00")

func _lgraphqlAddprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddprojectGraphql,
		"_lgraphql/addProject.graphql",
	)
}

func _lgraphqlAddprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddprojecttobillinggroupGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x49\x2f\xca\x2f\x2d\xb0\x52\x70\x07\x51\x9e\x79\x05\xa5\x25\x8a\x3a\x20\xe1\x82\xa2\xfc\xac\xd4\xe4\x12\x2b\x85\x00\x08\x03\x22\xa5\xa9\x50\xcd\xa5\xa0\xa0\xa0\x90\x98\x92\x02\x15\x0f\xc9\x77\xca\xcc\xc9\xc9\xcc\x4b\x07\x1b\xa0\x91\x09\x52\x66\x05\x55\xa5\xa0\x00\x35\x1c\x62\x89\x0e\x54\x10\x6e\x34\xcc\x12\xb0\x78\xad\x26\x5c\x57\x66\x0a\x94\x91\x97\x98\x9b\x0a\x91\xe4\x02\x61\x40\x00\x00\x00\xff\xff\x98\x52\x69\x5c\xb9\x00\x00\x00")

func _lgraphqlAddprojecttobillinggroupGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddprojecttobillinggroupGraphql,
		"_lgraphql/addProjectToBillingGroup.graphql",
	)
}

func _lgraphqlAddprojecttobillinggroupGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddprojecttobillinggroupGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addProjectToBillingGroup.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddsshkeyGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\xc1\x0a\xc2\x30\x10\x44\xef\xf9\x8a\x11\x3c\xb4\xe0\x17\xf4\xee\xc9\xa3\xe2\x7d\x21\x41\x17\x9b\x58\xda\xe4\x50\x24\xff\x2e\xe9\xa6\x59\xc1\x53\xc2\x9b\x61\xdf\xf8\x14\x29\xf2\x3b\xa0\x33\xc0\x31\x90\x77\x03\xae\x71\xe6\xf0\x38\x9c\x0a\x79\xb9\xf5\x4e\x63\xfa\xa7\xb7\x75\x2a\x70\x79\x5e\xe4\x2f\x41\x5a\xdc\x7c\xf6\xc4\x63\xeb\xf7\xf8\x18\x00\x20\x6b\xa5\xdc\x71\x98\x52\x1c\x2a\x06\xc4\xb9\xa9\x2b\x51\x67\xd3\x6b\x22\xde\x7d\x41\xe5\x45\xab\x17\x01\x27\x13\x74\x4e\x4d\xf2\xf6\xe6\xbe\x55\xd9\xfe\xac\x30\x7b\x25\x9b\x6f\x00\x00\x00\xff\xff\xbf\xff\xb7\x89\x17\x01\x00\x00")

func _lgraphqlAddsshkeyGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddsshkeyGraphql,
		"_lgraphql/addSshKey.graphql",
	)
}

func _lgraphqlAddsshkeyGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddsshkeyGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addSshKey.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAdduserGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\x31\x0e\xc2\x30\x0c\x85\xe1\x3d\xa7\x78\x48\x0c\xad\xc4\x09\x7a\x08\x16\xc4\x01\x2c\x12\x90\xa5\xda\x45\xa9\x3b\x21\xee\x8e\xda\x3a\x69\x06\x06\x2f\xff\x1b\x3e\xcb\x62\x64\x3c\x29\xba\x00\x9c\x93\x10\x8f\x03\x6e\x96\x59\x5f\xa7\xcb\x9a\x9e\x9c\x67\xbb\x92\xa4\x92\xb7\x3a\xd2\x9f\xf8\x98\x44\x92\x5a\x69\x3d\x3e\x01\x00\x28\xc6\xfb\x9c\x72\xc7\xfa\x5e\x6c\xf0\x08\xb8\xb5\x9b\xde\x1a\xec\x80\x7d\x3b\xc8\xaa\xfb\x52\xdd\xf2\xc1\xd6\xbf\x7d\xa5\x38\xb6\xe6\xbe\x86\xf5\x7e\x01\x00\x00\xff\xff\xf7\xf8\x9f\x0d\xfe\x00\x00\x00")

func _lgraphqlAdduserGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAdduserGraphql,
		"_lgraphql/addUser.graphql",
	)
}

func _lgraphqlAdduserGraphql() (*asset, error) {
	bytes, err := _lgraphqlAdduserGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addUser.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddusertogroupGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xbd\x0a\xc3\x30\x0c\x84\xf7\x3c\xc5\x05\x3a\x24\xd0\x27\xf0\x5e\xba\x75\xe8\xcf\x03\x18\x6c\x82\x20\xb6\x83\x62\x4f\x21\xef\x5e\xa4\xb8\xf1\xd2\x41\x20\xee\x8e\xfb\x2e\x94\x6c\x33\xa5\x88\xa1\x03\x2e\x65\xf5\x7c\x0b\x96\x66\x83\x57\x66\x8a\x53\x7f\x15\x79\xe2\x54\x96\x87\x0d\xfe\x8f\xfc\x4c\xb3\x37\xb8\xff\xde\x7e\xc4\xd6\x01\x80\x75\xee\xb3\x7a\x7e\x27\xb5\x06\x8a\x4b\xc9\xa6\x7a\x80\x80\x0c\x36\xf8\x03\xd6\xc0\xd8\x6b\x42\xcb\x25\x12\x95\xdb\x36\x9c\x09\x56\x72\x5b\xa1\xf2\x3e\x9e\x0c\x72\xf5\x91\x86\xc3\xec\xe4\xbe\x01\x00\x00\xff\xff\x00\xb0\x61\x57\xf3\x00\x00\x00")

func _lgraphqlAddusertogroupGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddusertogroupGraphql,
		"_lgraphql/addUserToGroup.graphql",
	)
}

func _lgraphqlAddusertogroupGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddusertogroupGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addUserToGroup.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\x4d\x0e\x82\x30\x10\x85\xf7\x9c\xe2\x91\xb8\x80\x84\x13\xb0\x74\xe7\x86\x98\xe8\x05\x08\x1d\xb5\x06\xa6\xb5\x0c\x26\x0d\xf1\xee\x86\x1f\x5b\x70\xd1\xb4\xdf\xd7\xce\xeb\x7b\x0d\xe4\x3c\xb2\x04\x38\x70\xdd\x51\x89\x8b\x38\xcd\xf7\xb4\x98\x8c\x75\xe6\x49\x8d\x94\x38\xb1\xa4\x39\xc6\x04\x00\x88\xdf\xda\x19\xee\x88\xe5\xe8\xab\xba\xa3\x6c\xd6\xc0\x32\x3f\xc7\x14\xab\x0a\x01\xbf\xa8\x7c\xbd\x18\xd7\x1d\xd0\x2a\x1c\xa7\xc9\x00\xce\x0c\xf2\x47\x7d\x40\x45\xb6\x35\xfe\xea\x6d\x7c\xb1\xa9\xb5\xf3\xc6\x12\xf7\x0f\x7d\x93\xf3\xd2\xa0\xda\x7e\x32\x58\x55\x0b\xc5\x06\x8d\xa3\x1d\x2b\x6a\x29\xf2\x27\x99\xd6\x37\x00\x00\xff\xff\xd5\xce\x35\x6e\x32\x01\x00\x00")

func _lgraphqlEnvironmentbynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentbynameGraphql,
		"_lgraphql/environmentByName.graphql",
	)
}

func _lgraphqlEnvironmentbynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentbynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environmentByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlMeGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x50\x50\xc8\x4d\x85\x32\x40\x20\x33\x05\xce\x4c\xcd\x4d\xcc\xcc\x81\xf3\xd2\x32\x8b\x8a\x4b\xfc\x12\x73\x53\xe1\x22\x39\x89\x48\x02\xb5\x5c\xb5\x80\x00\x00\x00\xff\xff\xd4\x0b\x94\x75\x54\x00\x00\x00")

func _lgraphqlMeGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlMeGraphql,
		"_lgraphql/me.graphql",
	)
}

func _lgraphqlMeGraphql() (*asset, error) {
	bytes, err := _lgraphqlMeGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/me.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\x4f\x6f\xdb\x3e\x0c\xbd\xf7\x53\xf0\x57\xfc\x0e\xdd\xc5\xd8\x76\xdc\x6d\xed\x8a\x6e\xe8\x96\x05\x49\x96\x6b\xc1\xc8\x74\xac\x45\x96\x1c\x8a\x76\x61\x04\xf9\xee\x83\x13\xd4\x96\xff\x24\x0d\xd0\x61\xba\x84\x79\x8f\x26\x9f\x48\x89\xda\x16\xc4\x15\xdc\x5c\x01\xfc\x6f\x31\xa3\x4f\x30\x17\xd6\x76\xfd\xdf\x3b\xd8\x5d\x01\x00\xe4\xec\x7e\x93\x92\xdb\x6a\x82\x19\xdd\x1c\x20\x80\xa3\xe7\xe1\x83\x17\xbf\x7a\xe9\xb8\x31\x6b\xaa\xf9\x83\x85\xb8\x6f\xb1\x69\x81\x15\xa3\x55\x29\xf9\x06\xc8\x0b\x63\x98\xb6\x05\x79\x09\x40\xd6\x25\x0a\x3d\x52\x15\x40\x2e\x2e\x94\x68\x67\xef\x6d\xa9\xd9\xd9\x8c\xac\xb4\x79\x94\xe8\x92\xe6\x95\x17\xca\xfc\x17\xca\x8d\xab\xc6\xb9\x05\xfa\xcd\x38\x33\xa3\xcc\x95\x34\xce\x4d\xd9\x65\x4e\x5a\xd2\x8b\x63\x5c\xd3\x1d\x1a\xd5\x60\x2e\x27\xeb\x53\x9d\xc8\xf4\x58\xb7\x29\x8a\x10\xdb\x86\x8f\xa9\x24\xe3\xf2\x5a\x76\xb0\x03\xff\x5d\x67\xba\xdd\xc7\x5a\xcb\x2f\x36\xa7\xcb\xb7\x66\x57\xe4\xbe\xad\x3b\x40\x14\x45\xe0\x2c\x3c\xd4\x04\x84\x04\xc0\xd3\x93\x54\x39\x75\xfa\x31\x68\x50\xbd\x32\xca\x56\xc4\xbe\xfb\x31\x40\xe1\x89\xfb\x18\x00\x65\xa8\xcd\x00\xf5\x3e\x7d\xa4\x6a\x10\x62\x34\xdd\x71\x6d\xa8\x5a\x54\xf9\x09\x66\x89\xa6\x18\x52\xfb\x01\x92\x68\xf6\x32\x19\x4b\x60\x70\x94\xe8\x87\x60\x67\xba\x2e\xa1\x43\x68\xef\x76\xa0\x13\xc0\x5c\x2f\x89\x1f\x98\x50\x88\x17\x29\xda\x9f\x7c\xbf\x2d\xd0\x40\x04\xd7\x1f\xa2\x8f\xd1\xfb\x6b\xd8\xef\x87\xcd\xb9\xd5\xc6\x68\xbb\x7e\x4b\x8f\x54\xc1\x4c\x56\x55\x1d\x70\x75\x8c\x3b\x77\x89\x3c\x23\xd3\x69\xe9\x64\xe3\x50\x58\x6b\x59\x27\x3a\xd1\x0a\xeb\x7b\xe5\x61\xe4\x5c\x4d\x02\x87\xb9\x41\xb5\xb9\x4c\xff\x33\xad\x52\xe7\x36\xaf\xec\x29\x45\x6b\xc9\x9c\x90\x3d\x22\x60\xe6\xd4\x86\xe4\x2e\x45\xf9\x67\x2a\xde\xd4\xf7\x50\xfb\x7d\x7d\x6f\x2e\x93\x7d\xb8\x62\x9f\xe3\x98\xc9\xfb\x73\xda\x5f\xc9\xf8\x43\x2b\x76\xde\x25\xb2\x20\xcc\xfc\x5f\xab\xd8\xa5\x47\xab\x99\x87\x61\xe2\xe0\x9d\x68\x3d\xc9\x96\x4b\x64\x8d\x2b\x43\x5d\x99\xbd\xcc\x5e\xb9\xce\xc0\x28\x3b\x53\xa2\x13\xaf\x19\xaf\xe7\xe2\xc5\x87\x77\xa2\x37\x85\x8e\xe0\x2d\x7a\x9a\x51\x32\xc0\xbf\x12\xc6\x63\xf8\x42\x4b\x67\x90\x04\x12\x7a\xf1\xfb\xcf\x44\x6f\x48\x0d\x26\xfe\x99\xf2\x8c\x1c\xe6\x7e\x89\xfa\x45\x0a\xcb\xf4\x62\xd5\xbf\xfb\xab\x3f\x01\x00\x00\xff\xff\x88\xc1\x98\xb6\x0e\x08\x00\x00")

func _lgraphqlProjectbynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectbynameGraphql,
		"_lgraphql/projectByName.graphql",
	)
}

func _lgraphqlProjectbynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectbynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projectByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"_lgraphql/addBillingGroup.graphql":               _lgraphqlAddbillinggroupGraphql,
	"_lgraphql/addEnvVariable.graphql":                _lgraphqlAddenvvariableGraphql,
	"_lgraphql/addGroup.graphql":                      _lgraphqlAddgroupGraphql,
	"_lgraphql/addGroupsToProject.graphql":            _lgraphqlAddgroupstoprojectGraphql,
	"_lgraphql/addNotificationEmail.graphql":          _lgraphqlAddnotificationemailGraphql,
	"_lgraphql/addNotificationMicrosoftTeams.graphql": _lgraphqlAddnotificationmicrosoftteamsGraphql,
	"_lgraphql/addNotificationRocketChat.graphql":     _lgraphqlAddnotificationrocketchatGraphql,
	"_lgraphql/addNotificationSlack.graphql":          _lgraphqlAddnotificationslackGraphql,
	"_lgraphql/addNotificationToProject.graphql":      _lgraphqlAddnotificationtoprojectGraphql,
	"_lgraphql/addOrUpdateEnvironment.graphql":        _lgraphqlAddorupdateenvironmentGraphql,
	"_lgraphql/addProject.graphql":                    _lgraphqlAddprojectGraphql,
	"_lgraphql/addProjectToBillingGroup.graphql":      _lgraphqlAddprojecttobillinggroupGraphql,
	"_lgraphql/addSshKey.graphql":                     _lgraphqlAddsshkeyGraphql,
	"_lgraphql/addUser.graphql":                       _lgraphqlAdduserGraphql,
	"_lgraphql/addUserToGroup.graphql":                _lgraphqlAddusertogroupGraphql,
	"_lgraphql/environmentByName.graphql":             _lgraphqlEnvironmentbynameGraphql,
	"_lgraphql/me.graphql":                            _lgraphqlMeGraphql,
	"_lgraphql/projectByName.graphql":                 _lgraphqlProjectbynameGraphql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"_lgraphql": &bintree{nil, map[string]*bintree{
		"addBillingGroup.graphql":               &bintree{_lgraphqlAddbillinggroupGraphql, map[string]*bintree{}},
		"addEnvVariable.graphql":                &bintree{_lgraphqlAddenvvariableGraphql, map[string]*bintree{}},
		"addGroup.graphql":                      &bintree{_lgraphqlAddgroupGraphql, map[string]*bintree{}},
		"addGroupsToProject.graphql":            &bintree{_lgraphqlAddgroupstoprojectGraphql, map[string]*bintree{}},
		"addNotificationEmail.graphql":          &bintree{_lgraphqlAddnotificationemailGraphql, map[string]*bintree{}},
		"addNotificationMicrosoftTeams.graphql": &bintree{_lgraphqlAddnotificationmicrosoftteamsGraphql, map[string]*bintree{}},
		"addNotificationRocketChat.graphql":     &bintree{_lgraphqlAddnotificationrocketchatGraphql, map[string]*bintree{}},
		"addNotificationSlack.graphql":          &bintree{_lgraphqlAddnotificationslackGraphql, map[string]*bintree{}},
		"addNotificationToProject.graphql":      &bintree{_lgraphqlAddnotificationtoprojectGraphql, map[string]*bintree{}},
		"addOrUpdateEnvironment.graphql":        &bintree{_lgraphqlAddorupdateenvironmentGraphql, map[string]*bintree{}},
		"addProject.graphql":                    &bintree{_lgraphqlAddprojectGraphql, map[string]*bintree{}},
		"addProjectToBillingGroup.graphql":      &bintree{_lgraphqlAddprojecttobillinggroupGraphql, map[string]*bintree{}},
		"addSshKey.graphql":                     &bintree{_lgraphqlAddsshkeyGraphql, map[string]*bintree{}},
		"addUser.graphql":                       &bintree{_lgraphqlAdduserGraphql, map[string]*bintree{}},
		"addUserToGroup.graphql":                &bintree{_lgraphqlAddusertogroupGraphql, map[string]*bintree{}},
		"environmentByName.graphql":             &bintree{_lgraphqlEnvironmentbynameGraphql, map[string]*bintree{}},
		"me.graphql":                            &bintree{_lgraphqlMeGraphql, map[string]*bintree{}},
		"projectByName.graphql":                 &bintree{_lgraphqlProjectbynameGraphql, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
