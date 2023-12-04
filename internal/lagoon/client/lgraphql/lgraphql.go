// Code generated for package lgraphql by go-bindata DO NOT EDIT. (@generated)
// sources:
// _lgraphql/addDeployTarget.graphql
// _lgraphql/addEnvVariable.graphql
// _lgraphql/addGroup.graphql
// _lgraphql/addGroupsToProject.graphql
// _lgraphql/addOrUpdateEnvironment.graphql
// _lgraphql/addProject.graphql
// _lgraphql/addRestore.graphql
// _lgraphql/addSshKey.graphql
// _lgraphql/addUser.graphql
// _lgraphql/addUserToGroup.graphql
// _lgraphql/backupsForEnvironmentByName.graphql
// _lgraphql/deleteDeployTarget.graphql
// _lgraphql/deleteDeployTargetConfig.graphql
// _lgraphql/deployEnvironmentBranch.graphql
// _lgraphql/deployEnvironmentLatest.graphql
// _lgraphql/deployEnvironmentPromote.graphql
// _lgraphql/deployEnvironmentPullrequest.graphql
// _lgraphql/deployTargetConfigsByProjectId.graphql
// _lgraphql/environmentByName.graphql
// _lgraphql/lagoonSchema.graphql
// _lgraphql/lagoonVersion.graphql
// _lgraphql/listDeployTargets.graphql
// _lgraphql/me.graphql
// _lgraphql/minimalProjectByName.graphql
// _lgraphql/projectByName.graphql
// _lgraphql/projectByNameMetadata.graphql
// _lgraphql/sshEndpointsByProject.graphql
// _lgraphql/updateDeployTarget.graphql
// _lgraphql/updateDeployTargetConfig.graphql
// _lgraphql/variables/addOrUpdateEnvVariableByName.graphql
// _lgraphql/variables/deleteEnvVariableByName.graphql
// _lgraphql/variables/getEnvVariablesByProjectEnvironmentName.graphql
// _lgraphql/notifications/addNotificationEmail.graphql
// _lgraphql/notifications/addNotificationMicrosoftTeams.graphql
// _lgraphql/notifications/addNotificationRocketChat.graphql
// _lgraphql/notifications/addNotificationSlack.graphql
// _lgraphql/notifications/addNotificationToProject.graphql
// _lgraphql/notifications/addNotificationWebhook.graphql
// _lgraphql/notifications/deleteNotificationEmail.graphql
// _lgraphql/notifications/deleteNotificationMicrosoftTeams.graphql
// _lgraphql/notifications/deleteNotificationRocketChat.graphql
// _lgraphql/notifications/deleteNotificationSlack.graphql
// _lgraphql/notifications/deleteNotificationWebhook.graphql
// _lgraphql/notifications/listAllNotificationEmail.graphql
// _lgraphql/notifications/listAllNotificationMicrosoftTeams.graphql
// _lgraphql/notifications/listAllNotificationRocketChat.graphql
// _lgraphql/notifications/listAllNotificationSlack.graphql
// _lgraphql/notifications/listAllNotificationWebhook.graphql
// _lgraphql/notifications/projectNotificationEmail.graphql
// _lgraphql/notifications/projectNotificationMicrosoftTeams.graphql
// _lgraphql/notifications/projectNotificationRocketChat.graphql
// _lgraphql/notifications/projectNotificationSlack.graphql
// _lgraphql/notifications/projectNotificationWebhook.graphql
// _lgraphql/notifications/removeNotificationFromProject.graphql
// _lgraphql/notifications/updateNotificationEmail.graphql
// _lgraphql/notifications/updateNotificationMicrosoftTeams.graphql
// _lgraphql/notifications/updateNotificationRocketChat.graphql
// _lgraphql/notifications/updateNotificationSlack.graphql
// _lgraphql/notifications/updateNotificationWebhook.graphql
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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

// Mode return file modify time
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

var __lgraphqlAdddeploytargetGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xd1\xcf\x4e\xc3\x30\x0c\x06\xf0\x7b\x9f\xc2\x48\x1c\x36\x69\x4f\xd0\x2b\x1c\x80\xc3\x56\x31\x78\x80\x40\xdc\x2e\x22\x8d\x2b\xd7\x45\x9a\x10\xef\x8e\xfa\x27\x55\xbc\x94\x43\x0f\xfd\xbe\x44\xb2\x7f\x69\x07\x31\xe2\x28\xc0\xae\x00\xb8\x0f\xa6\xc5\x12\xce\xc2\x2e\x34\x77\x87\x31\xf9\xa4\xd0\x93\xc7\x77\xf6\x3a\x17\xfa\xc2\x10\xa3\x29\x61\x1a\x04\xb9\x32\x22\xc8\xba\xe9\xfb\xcb\x13\xf5\x72\x9b\x55\xc4\x3a\x6b\x29\x38\xa1\xf1\xf7\x81\x42\xed\x9a\x12\x5e\xce\xa7\xe3\x54\x7d\x0c\xce\xdb\xe7\xd6\x34\xa8\x6e\xd4\xec\x30\x58\x7f\x3d\x26\x73\xcf\x63\x7b\x1a\x6c\xc5\xf4\xed\x2c\x72\xde\xbc\x62\xe3\x68\x1d\x72\x0f\x3f\x05\x00\x80\xb1\xf6\x11\x3b\x4f\xd7\x37\xc3\x0d\x4a\x39\x06\xa7\x0e\x43\x7f\x71\xb5\xec\x5c\xe8\x06\x29\x97\xa3\x00\x33\xd5\x24\x76\x58\xa2\xd4\x2a\x81\x8b\xf5\x42\x36\xd3\x2d\xd9\x0d\x9a\x46\x8c\x17\x57\xbf\x28\x99\x14\x33\x62\xe4\x8c\x45\xca\x95\xd8\xc5\x3a\x87\xce\xec\xe3\x51\x2d\xac\xc0\xd7\xb5\xb5\xb5\xb6\x57\x87\x22\x7b\xfa\x08\x53\xff\xbb\x5f\x5d\x9d\x4d\x80\xe3\x65\x46\x23\x68\x53\xc7\x8c\x5c\x5b\x69\xa0\xad\x41\xf3\xb9\x36\x36\xce\x38\xb7\x9e\xed\x1f\xd4\x79\xb1\x62\xfc\xfe\x02\x00\x00\xff\xff\x0e\x41\x38\x1f\x65\x03\x00\x00")

func _lgraphqlAdddeploytargetGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAdddeploytargetGraphql,
		"_lgraphql/addDeployTarget.graphql",
	)
}

func _lgraphqlAdddeploytargetGraphql() (*asset, error) {
	bytes, err := _lgraphqlAdddeploytargetGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addDeployTarget.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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

var __lgraphqlAddprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x4b\x6a\xc3\x40\x0c\x86\xf7\x3e\x85\x02\x5d\x24\x90\x13\x78\x5b\xba\x08\xed\x22\x50\x7a\x80\x89\x47\x71\xa6\x8c\x35\xae\xac\x31\x94\xd2\xbb\x17\x3f\xe6\x95\x26\x59\x78\xf3\xff\x92\x2d\x7f\x5f\xe7\x45\x89\x71\x04\xdb\x0a\xe0\x89\x54\x87\x35\xbc\x0b\x1b\x6a\x37\xfb\x29\x69\x8d\x7c\xb0\x2d\xb3\xc1\x9f\xce\xce\x6a\xe4\x10\xcf\xa9\xeb\x91\x86\x8b\x39\x4b\x0d\x07\x92\x4d\x99\x1d\xd9\x7d\x62\x23\x47\x25\x82\x4c\xc5\xde\x89\x15\x35\x17\x1c\x8a\xb0\xf7\xd6\x32\x7e\x79\x1c\xe4\xaa\x60\xa7\x7d\x33\x5d\xfc\x42\xa3\x61\x47\x1d\x92\x94\xe7\x29\x2f\xee\xa0\x2d\xce\x77\x2c\x07\x8b\x63\xd5\xe2\xb3\xb2\x4d\x0a\x35\x8e\x68\x5d\x3f\xed\x67\xaf\x1a\xde\x4c\x67\x24\x4d\xf5\x6c\x46\x25\xf8\x8a\xdf\xe1\x23\x3b\xf8\xa9\x00\x00\x94\xd6\xeb\x5f\x6d\x0d\xf5\x5e\xea\x35\x07\x58\x28\xce\x30\xf7\x6b\x14\x30\xae\x3c\x43\x9c\x91\x4c\x54\x43\x99\x01\x4d\x20\xff\x95\xd7\x64\xef\x31\x0f\x8b\x09\x77\x24\x1f\xaa\x12\x7a\xe1\x20\x8e\xdc\xc6\x7f\x5b\x4b\x58\x4a\x42\xa2\x9b\x08\x20\x37\x93\x7b\x0a\x03\x8f\x2d\x3d\x94\x98\x6e\x4e\x0a\x33\x9f\x73\xfb\xbb\x8b\xd6\x8c\xce\xf4\x2d\x65\x35\x3d\x7f\x01\x00\x00\xff\xff\x70\x99\xe2\x23\x22\x03\x00\x00")

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

var __lgraphqlAddrestoreGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x49\x4a\x4c\xce\x2e\x2d\xc8\x4c\xb1\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x4c\x49\x09\x4a\x2d\x2e\xc9\x2f\x4a\xd5\xc8\xcc\x2b\x28\x2d\xb1\x52\xa8\x86\xa8\xf5\x4c\xb1\x42\x68\xab\x85\x29\x57\x50\xc8\x4c\x01\x33\x6a\xb9\x40\x18\x10\x00\x00\xff\xff\xdc\x18\x6c\xe7\x65\x00\x00\x00")

func _lgraphqlAddrestoreGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddrestoreGraphql,
		"_lgraphql/addRestore.graphql",
	)
}

func _lgraphqlAddrestoreGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddrestoreGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addRestore.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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

var __lgraphqlBackupsforenvironmentbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x4d\x6e\x84\x30\x0c\x85\xf7\x39\xc5\x43\xea\x02\x24\x4e\xc0\xb2\x3b\xa4\xaa\x9b\x9e\x20\x0d\x56\x95\xb6\x24\x8c\xe3\x8c\x84\x10\x77\x1f\xf1\x33\x99\xc0\xcc\xca\xce\x67\xbf\xcf\xb9\x44\xe2\x11\xa5\x02\xde\x9c\xee\xa9\xc1\x97\xb0\x75\x3f\x45\xbd\x90\x81\xfd\x2f\x19\x69\xd0\x3a\x29\x2a\x4c\x0a\x00\xc8\x5d\x2d\x7b\xd7\x93\x93\xf7\xf1\x53\xf7\x54\xae\x18\xd8\xf2\xab\xa6\xde\x51\x12\xdc\x55\xd5\x3e\x98\xf6\x0a\xd8\x2e\xb5\x4b\x32\x3d\xbe\xb5\xf9\x8b\x43\x78\x2c\x1e\x56\x81\xe0\x23\x1b\xca\xc0\x16\x68\xf3\x1d\xc3\xa4\x85\x72\xd2\xd1\x3f\x1d\x09\x53\x10\xcf\x94\xdf\x39\x5d\x02\x82\x68\x89\xe1\x80\x9e\xd5\x2f\x7f\x90\xfc\x1f\xde\x68\xb1\xde\x65\xb3\x59\x9d\xbb\xa5\xce\xea\x16\x00\x00\xff\xff\x0c\x15\x4e\x54\x93\x01\x00\x00")

func _lgraphqlBackupsforenvironmentbynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlBackupsforenvironmentbynameGraphql,
		"_lgraphql/backupsForEnvironmentByName.graphql",
	)
}

func _lgraphqlBackupsforenvironmentbynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlBackupsforenvironmentbynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/backupsForEnvironmentByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeletedeploytargetGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xe4\xd2\x54\xa8\xe6\x52\x50\x48\x49\xcd\x49\x2d\x49\x75\x49\x2d\xc8\xc9\xaf\x0c\x49\x2c\x4a\x4f\x2d\xb1\x82\x8a\xf9\x17\xa4\xe6\x15\x67\x64\xa6\x95\x68\x64\xe6\x15\x94\x96\x58\x81\x95\x2b\x28\x40\x0c\x01\x9b\xc5\xa5\xa0\x50\xab\xc9\x55\xcb\x05\x08\x00\x00\xff\xff\x8e\x88\xf8\xcc\x66\x00\x00\x00")

func _lgraphqlDeletedeploytargetGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeletedeploytargetGraphql,
		"_lgraphql/deleteDeployTarget.graphql",
	)
}

func _lgraphqlDeletedeploytargetGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeletedeploytargetGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deleteDeployTarget.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeletedeploytargetconfigGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x50\x50\xc9\x4c\xb1\x52\xf0\xcc\x2b\x51\x84\xf0\x0a\x8a\xf2\xb3\x52\x93\x4b\xa0\x42\x9a\xd5\x60\xd1\x94\xd4\x9c\xd4\x92\x54\x97\xd4\x82\x9c\xfc\xca\x90\xc4\xa2\xf4\xd4\x12\xe7\xfc\xbc\xb4\xcc\x74\x8d\xcc\xbc\x82\xd2\x12\x2b\x88\x22\x10\x00\x99\xa5\x92\x99\x02\xe7\xc3\x4d\x83\x99\x0b\x96\xa9\xd5\xe4\xaa\x05\x04\x00\x00\xff\xff\x66\x8a\x2f\xa2\x86\x00\x00\x00")

func _lgraphqlDeletedeploytargetconfigGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeletedeploytargetconfigGraphql,
		"_lgraphql/deleteDeployTargetConfig.graphql",
	)
}

func _lgraphqlDeletedeploytargetconfigGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeletedeploytargetconfigGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deleteDeployTargetConfig.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeployenvironmentbranchGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x41\x0a\xc2\x30\x10\x45\xf7\x39\xc5\x2f\xb8\x68\xa1\x27\xc8\xb2\xe8\xd6\x85\x9e\x20\xd6\x51\x23\xed\xa4\x0c\x53\x41\xa4\x77\x97\x36\x89\x56\x70\x35\xcc\xfb\x9f\xcf\xeb\x47\x75\xea\x03\xa3\x34\xc0\x66\x90\x70\xa7\x56\x2d\x8e\x2a\x9e\xaf\x45\x8d\x99\x9e\xc4\x71\x7b\xfb\x0b\x0f\x74\xc9\xbc\x9e\xa9\x90\x8e\xc2\x5b\xa7\xce\xa2\x09\xa1\x23\xc7\x45\x85\x97\x01\x80\x33\x0d\x5d\x78\xee\xf8\xe1\x25\x70\x4f\xac\xcd\x32\x51\x7a\x1e\x46\xb5\xa9\x04\x64\x87\xfc\x03\xec\x7a\xb2\x1f\xb9\x84\xa7\x74\xa3\xc7\x3e\x56\xe2\xf3\x93\x2c\x86\x5f\xdb\x94\xad\x3d\x57\xd2\x26\x2f\x57\x66\x7a\x07\x00\x00\xff\xff\xb1\xcb\x5e\x73\x1a\x01\x00\x00")

func _lgraphqlDeployenvironmentbranchGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeployenvironmentbranchGraphql,
		"_lgraphql/deployEnvironmentBranch.graphql",
	)
}

func _lgraphqlDeployenvironmentbranchGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeployenvironmentbranchGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployEnvironmentBranch.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeployenvironmentlatestGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcd\x41\x0a\xc2\x40\x0c\x85\xe1\xfd\x9c\xe2\x15\xba\x68\xc1\x13\xcc\x52\x74\x21\x78\x89\x80\x59\x0c\xb4\x49\x19\xdf\x08\x22\xbd\xbb\x28\x03\xc6\x65\x78\xf9\xf9\xd6\x46\x61\x71\xc3\x94\x80\x51\xed\x51\xaa\xdb\xaa\xc6\x8c\xf3\xef\xb8\xd8\xd6\x38\x1c\x30\x56\x65\xab\x76\x12\x4a\xc6\xd1\x7d\x51\xb1\x61\xc6\x2b\x01\xc0\x4d\xb7\xc5\x9f\xa1\xba\x0a\xf5\xce\xa9\x7c\xe2\xdc\x9f\x80\x3f\x23\x8a\x7d\x8f\x44\xf0\xbe\xeb\x9e\x80\x39\xed\xef\x00\x00\x00\xff\xff\xf5\x45\x27\xcb\xb5\x00\x00\x00")

func _lgraphqlDeployenvironmentlatestGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeployenvironmentlatestGraphql,
		"_lgraphql/deployEnvironmentLatest.graphql",
	)
}

func _lgraphqlDeployenvironmentlatestGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeployenvironmentlatestGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployEnvironmentLatest.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeployenvironmentpromoteGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xc1\xaa\x83\x30\x10\x45\xf7\xf9\x8a\x2b\xbc\x85\xc2\xfb\x82\x2c\x1f\xaf\xfb\x42\xbf\x20\xe8\x50\x52\xcc\x8c\xc4\x49\xa1\x14\xff\xbd\xd8\x1a\x6b\xab\xae\x02\xb9\xe7\x0c\x33\x37\x24\x75\xea\x85\x51\x1a\xe0\xa7\x8b\x72\xa1\x5a\x2d\x4e\x1a\x3d\x9f\x8b\x5f\x8c\xbf\xbd\xa4\x58\xd3\x81\xaf\x3e\x0a\x07\xe2\xef\xbc\xa1\x5e\x3d\x3f\xc7\x6c\x42\x23\x13\x49\x53\xe4\x7f\xa7\xce\xe2\x4f\xa4\x25\xc7\x45\x85\xbb\x01\x80\x86\xba\x56\x6e\x0b\xf5\x18\x25\x88\x52\xe9\xb9\x4b\x6a\x5f\x10\xb0\x5e\x23\x27\x00\xbb\x40\x76\x63\xd3\x19\xc8\x97\xbd\x95\x59\x9a\xa2\x39\x18\xcc\xe7\xbb\x56\x37\xc5\x8c\xef\xb5\xb1\x53\xd3\x64\x2d\xfb\x59\x94\x65\xf2\xe4\xca\x0c\x8f\x00\x00\x00\xff\xff\xcd\xc9\x92\xe1\xab\x01\x00\x00")

func _lgraphqlDeployenvironmentpromoteGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeployenvironmentpromoteGraphql,
		"_lgraphql/deployEnvironmentPromote.graphql",
	)
}

func _lgraphqlDeployenvironmentpromoteGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeployenvironmentpromoteGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployEnvironmentPromote.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeployenvironmentpullrequestGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x41\x4e\xc3\x30\x10\x45\xf7\x39\xc5\x8f\xc4\xa2\x95\x7a\x02\x2f\x2b\x58\x74\x83\x2a\x38\xc1\xb4\x1d\xa8\x91\x33\x0e\xd3\x31\x12\x42\xbd\x3b\x4a\x33\x85\xc6\x51\x77\xc9\xff\x7f\x2c\xbd\xd7\x15\x23\x8b\x59\xb0\x68\x80\x87\x5e\xf3\x07\xef\x2d\x60\x3b\x7e\x6c\xa4\x2f\xd6\xae\x30\x74\x52\xba\x1d\x6b\xc0\x46\xac\x5d\x0d\x81\x45\x4b\x1c\xf0\x6a\x1a\xe5\x7d\x8c\x76\x74\xe2\xb5\x92\xec\x8f\xcf\xd4\xdd\xed\x5e\xf8\x6d\x5a\x1d\x99\x0e\xf7\xce\xfe\xbb\xd9\x99\xb2\x15\x95\x47\x32\x0a\x58\xe7\x9c\x98\xa4\x5d\xe2\xa7\x01\x80\x03\xf7\x29\x7f\x3f\xc9\x57\xd4\x2c\x1d\x8b\x6d\x4b\x4a\xca\x9f\x85\x4f\xb6\x88\x03\x56\xf0\x25\xf0\x47\x7d\xe5\xf7\xfc\x4a\xec\xe8\x9e\x3a\xf6\x88\xef\x59\xcd\x5d\x89\x98\xad\x2e\x28\x53\x23\xbe\xa9\x55\x54\x6e\x66\xab\xf1\xa5\xc9\xbf\x6f\x6e\xed\xdc\xa8\xba\xb4\xe7\x06\x58\x36\xe7\xdf\x00\x00\x00\xff\xff\x49\x97\xf5\x67\xfd\x01\x00\x00")

func _lgraphqlDeployenvironmentpullrequestGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeployenvironmentpullrequestGraphql,
		"_lgraphql/deployEnvironmentPullrequest.graphql",
	)
}

func _lgraphqlDeployenvironmentpullrequestGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeployenvironmentpullrequestGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployEnvironmentPullrequest.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetconfigsbyprojectidGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xc1\xca\xc2\x30\x10\x84\xef\x79\x8a\xfd\xe1\x3f\xe8\x2b\x78\xd4\x53\x2f\x52\xc4\x17\xa8\xcd\x34\x8d\xc4\x6c\xbb\x4d\x94\x20\x7d\x77\x69\x82\xa5\xe6\x90\x65\xbf\xd9\x19\x66\x8c\x90\x44\x3b\x45\xf4\x3f\x08\xdf\xd1\x86\x03\x55\x3e\xfc\xa9\xfd\x5b\x11\x69\x0c\x8e\xd3\xb5\x11\x83\x70\x62\xdf\x59\x33\x1d\x53\x5d\xee\x2a\xbd\xb8\x88\x56\xdb\x37\x40\x11\x65\x33\x91\xd5\x79\x6c\x53\x8a\xb0\x11\x97\xe7\x9b\x07\xd6\xa5\x13\x0b\xaf\x5d\x3a\x6f\x61\xeb\x38\xea\x5a\xf8\x69\x35\xe4\x97\x5e\x60\x2c\xfb\xcc\xe6\xfc\xdf\xa4\xf1\x6d\x8f\xa9\xb4\x8b\xce\x09\xc6\x88\x29\x14\xf0\x82\x35\xfd\xd2\x71\x56\xf3\x27\x00\x00\xff\xff\xba\x92\xa1\x8b\xfd\x00\x00\x00")

func _lgraphqlDeploytargetconfigsbyprojectidGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetconfigsbyprojectidGraphql,
		"_lgraphql/deployTargetConfigsByProjectId.graphql",
	)
}

func _lgraphqlDeploytargetconfigsbyprojectidGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetconfigsbyprojectidGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployTargetConfigsByProjectId.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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

var __lgraphqlLagoonschemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x88\x8f\x2f\x4e\xce\x48\xcd\x4d\x04\x73\x14\x14\x4a\x2a\x0b\x52\x8b\xa1\x6c\x05\x85\xec\xcc\xbc\x14\x28\x33\x2f\x31\x37\x15\xca\x4c\xcb\x4c\xcd\x49\x29\xd6\xc8\xcc\x4b\xce\x29\x4d\x49\x75\x49\x2d\x28\x4a\x4d\x4e\x2c\x49\x4d\xb1\x2a\x29\x2a\x4d\xd5\x84\x6b\x46\xd1\x53\xcb\x05\x23\x6b\xb9\x6a\x01\x01\x00\x00\xff\xff\x29\x07\x39\xef\x7e\x00\x00\x00")

func _lgraphqlLagoonschemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlLagoonschemaGraphql,
		"_lgraphql/lagoonSchema.graphql",
	)
}

func _lgraphqlLagoonschemaGraphql() (*asset, error) {
	bytes, err := _lgraphqlLagoonschemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/lagoonSchema.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlLagoonversionGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\xc8\x49\x4c\xcf\xcf\xcf\x0b\x4b\x2d\x2a\xce\xcc\xcf\xe3\xaa\x05\x04\x00\x00\xff\xff\x42\xb4\x77\x45\x19\x00\x00\x00")

func _lgraphqlLagoonversionGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlLagoonversionGraphql,
		"_lgraphql/lagoonVersion.graphql",
	)
}

func _lgraphqlLagoonversionGraphql() (*asset, error) {
	bytes, err := _lgraphqlLagoonversionGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/lagoonVersion.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlListdeploytargetsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8c\xd1\x4a\xc4\x40\x0c\x45\xdf\xfb\x15\xf9\x0e\x5f\xf5\x41\x11\xa4\x88\x7e\xc0\xb4\x73\x3b\x06\xd3\x44\x33\x99\x85\xb2\xf4\xdf\x97\xed\x50\xf6\x29\xe7\x5e\xee\xc9\x7f\x83\x6f\x74\x1d\x88\x84\x6b\xbc\xe0\x4f\x6c\xfb\x4a\x5e\x10\xf5\x89\x92\xc8\x7b\x9b\xe0\x8a\x40\xbd\x6f\x88\x38\x1f\x47\xd3\x8a\x03\x66\x47\x0a\xf4\x72\x36\xad\x26\xf8\x76\xe9\x51\xac\xe5\x4f\x14\x36\x7d\xe4\xd1\xed\xc2\x19\x7e\x34\x8b\x33\x34\xcb\xf6\x71\x7e\x0b\xfb\x45\x1f\xd7\xfa\xf3\x6a\x35\x4e\x1e\xcd\x3b\x4f\x8d\x25\xbf\xad\xa9\x74\xc1\xad\x05\x7c\x4c\x11\xf0\x2e\xae\xa6\x1c\xe6\xac\xe5\xd9\x74\xe1\x32\x10\xed\xc3\x7e\x0b\x00\x00\xff\xff\x3e\x5c\x37\x43\xe7\x00\x00\x00")

func _lgraphqlListdeploytargetsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlListdeploytargetsGraphql,
		"_lgraphql/listDeployTargets.graphql",
	)
}

func _lgraphqlListdeploytargetsGraphql() (*asset, error) {
	bytes, err := _lgraphqlListdeploytargetsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/listDeployTargets.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlMeGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x50\x50\xc8\x4d\x85\x32\x40\x20\x33\x05\xce\x4c\xcd\x4d\xcc\xcc\x81\xf3\xd2\x32\x8b\x8a\x4b\xfc\x12\x73\x53\xe1\x22\x39\x89\x68\x02\xc5\xc5\x19\xde\xa9\x95\xc5\x48\xa6\xa1\x99\x08\x02\x79\xc8\x3a\x40\x20\x3b\xb5\x32\xa4\xb2\x00\x43\x2c\x2c\x31\xa7\x14\x43\xd0\x2d\x33\x2f\x3d\xb5\xa8\xa0\x28\x33\xaf\x04\x45\x2a\xb9\x28\x35\xb1\x24\x15\x61\x51\x2d\x17\x84\xac\x05\x04\x00\x00\xff\xff\x19\x81\x0c\x89\xe8\x00\x00\x00")

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

var __lgraphqlMinimalprojectbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x41\x4a\xc6\x30\x10\x85\xf7\x39\xc5\x08\x2e\xea\x15\x5c\x0a\x2e\x04\x91\x82\x78\x80\xd8\x8c\xed\x48\x32\x49\x27\x93\x42\x91\xde\x5d\xda\x62\x52\x17\x7f\x36\x49\xbe\x37\x2f\x79\x6f\x2e\x28\x2b\x74\x06\xe0\x9e\x6d\xc0\x47\x78\x57\x21\x1e\xef\x1e\xe0\xc7\x00\x00\x24\x89\xdf\x38\xe8\xd3\xfa\x66\x03\x76\x07\x02\x38\x27\x0f\xc3\xdf\xdc\xbe\xc8\xd5\xe3\x2e\xd5\x8b\x2d\x1a\x5f\x9c\x6f\xe0\x53\x2c\x0f\x13\xe6\x0a\x52\xf1\x5e\x70\x2e\x98\xf5\x02\x25\xba\x32\x28\x45\x7e\xe6\x85\x24\x72\x40\xd6\xaa\xc6\x84\x9c\x27\xfa\xd2\xfe\x4c\xd8\x5b\x55\x14\xae\xba\xc3\x05\x7d\x4c\xbb\xe7\x62\xcf\xaf\x14\xa8\x3d\x32\x92\x7e\x88\xbf\x1d\xb4\x7e\xd2\x4a\xfe\xab\xb9\x99\xb6\x6f\xe6\x37\x00\x00\xff\xff\x43\xb8\x3d\x3c\x4c\x01\x00\x00")

func _lgraphqlMinimalprojectbynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlMinimalprojectbynameGraphql,
		"_lgraphql/minimalProjectByName.graphql",
	)
}

func _lgraphqlMinimalprojectbynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlMinimalprojectbynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/minimalProjectByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\xc1\x6e\xdb\x30\x0c\xbd\xe7\x2b\xb8\x62\x87\xee\x62\x6c\x3b\xee\xb6\x06\x41\x37\x74\xcb\x8a\x36\xcd\xb5\x60\x6c\xda\xd6\x22\x4b\x8e\x44\xa7\x30\x02\xff\xfb\xe0\x24\xb3\x25\x5b\x4e\x7b\x18\xe6\x4b\x9c\x47\xe9\xe9\xf1\x91\xa2\x77\x15\x99\x1a\xae\x67\x00\xef\x15\x16\xf4\x05\x1e\xd9\x08\x95\xbd\xfb\x00\x87\x19\x00\x40\x69\xf4\x6f\x8a\xf9\xa6\x5e\x62\x41\xd7\x47\x08\xe0\xb4\xf2\xb8\xe1\xef\xba\xf6\x11\x49\xf7\xda\x86\xba\x3f\x58\xb1\xfe\x9e\xc8\x1e\xd8\x18\x54\x71\x4e\xb6\x03\xca\x4a\x4a\x43\xbb\x8a\x2c\x3b\xa0\x11\x7b\x64\xba\xa3\xda\x81\x74\x52\xc5\x2c\xb4\x5a\xa8\xbd\x30\x5a\x15\xa4\xb8\x8b\x5a\xd6\x06\x33\x9a\xa3\x8c\x3b\x4c\x97\xa4\x6c\x2e\x52\xbe\x3f\x25\x72\x8f\xcc\x64\x54\x17\x4f\x68\x4f\x52\x97\x2d\x8f\x43\x69\x7f\x88\x42\xf4\xc4\x99\xe0\x27\x23\xa7\xf3\x49\x31\x66\xfb\x24\x5c\x99\x1b\x49\x85\x0b\x25\x54\x4a\x5d\xaf\xd0\x64\xc4\x73\xad\x52\x91\x59\xc7\xb9\x80\x25\xfe\x16\x77\xa9\x67\xf4\xc8\xec\xf6\x61\xbd\x25\xe5\x20\xcd\x6c\x62\xef\xb9\xba\x3e\xfb\x28\xbd\x80\x05\xc1\x63\x2f\x57\x07\x26\x2a\xe4\xcb\x0b\x76\x02\xc0\x0b\x89\x2c\xef\xb9\xfa\x1d\x99\xd1\x55\x69\xdd\x04\xa2\x28\x02\xad\xe0\xb6\x0d\x80\x9f\xd9\xf3\x33\xd7\x25\x8d\x74\x8f\x80\x82\x8a\x0d\x19\xeb\x6f\x06\xa8\x2c\x99\x21\x06\x40\x05\x0a\x39\x42\xad\xcd\xef\xa8\x1e\x51\x04\x8f\x3b\x3d\x5b\xaa\x57\x75\x39\x11\x59\xa3\xac\xc6\xa1\x66\x84\xa4\xc2\x58\x5e\x86\x0e\x90\x18\x0c\x0c\x29\x8c\x1e\x14\xbe\x09\x16\xaa\x7f\x53\x9a\x45\x2a\x62\x6c\xcb\xee\xb7\xf4\xb9\x10\x4b\x67\xc1\xa3\xc4\x78\xfb\xb6\xa2\xbc\xd0\x26\xd7\x7a\x7b\xb9\x50\x71\x8e\x4a\x91\x9c\xe8\xa5\x80\x80\x07\x1d\x6f\x89\xe7\x39\xf2\x7f\x53\x71\x38\x80\x48\x01\x4b\xb1\x26\x73\x6b\x08\x99\xcc\x2a\x47\xf5\xcb\x2c\x76\x15\x4a\x88\xe0\xea\x53\xf4\x39\xfa\x78\x05\xcd\x2b\xda\x17\x6d\xa3\xbd\x4d\xf6\xb1\x27\xbf\x26\x89\x21\x6b\x2f\x69\x7f\xe5\xc4\x9f\x22\x36\xda\xea\x94\x57\x84\x85\xfd\x67\x8e\x0d\xdc\x21\x95\xb8\xb9\x37\xe3\xd9\x7d\x08\xcf\xaf\x7e\x25\xa9\xfd\x1a\x8d\xc0\x8d\x24\x5f\xe6\xe0\x64\x1b\x6b\xef\x86\xed\xbd\x6b\xe5\xf1\x75\x9f\x82\x4b\x7c\xe7\x11\xed\x5f\xdb\x13\x78\x83\x96\x1e\x28\x1d\xe1\xdf\x08\x93\x10\xbe\x12\xec\xdd\x3c\x47\xc2\x80\x7f\xf8\x49\x1b\xdc\xea\xc0\xf8\x9e\xb4\x27\xd0\xcc\x43\x8b\x86\x26\x85\xa6\x40\xfb\xdb\xcc\xfe\x04\x00\x00\xff\xff\x87\x73\x32\xb1\x4b\x08\x00\x00")

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

var __lgraphqlProjectbynamemetadataGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x71\xaa\xf4\x4b\xcc\x4d\xd5\x00\x0b\x29\x28\x40\x54\x82\x35\xc0\xd4\x81\x40\x66\x0a\x9c\x09\x92\x82\x73\x72\x53\x4b\x12\x53\x12\x4b\x12\xa1\x02\xb5\x5c\x20\x0c\x08\x00\x00\xff\xff\xd4\x75\x13\x1d\x79\x00\x00\x00")

func _lgraphqlProjectbynamemetadataGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectbynamemetadataGraphql,
		"_lgraphql/projectByNameMetadata.graphql",
	)
}

func _lgraphqlProjectbynamemetadataGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectbynamemetadataGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projectByNameMetadata.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlSshendpointsbyprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x71\xaa\xf4\x4b\xcc\x4d\xd5\x80\x28\x00\xab\x83\x48\x2b\x28\x64\xa6\x80\x29\x90\x10\x98\x91\x9a\x57\x96\x59\x94\x9f\x97\x9b\x9a\x57\x52\x0c\x55\x02\x57\x84\xa4\x4c\x41\x21\xbf\x20\x35\xaf\x38\x23\x33\xad\x24\x00\x62\x85\x1f\x16\x29\x98\x7e\x24\x13\x14\x14\x8a\x8b\x33\x3c\xf2\x8b\x4b\x90\xf9\x01\xf9\x45\x30\x7e\x2d\x17\x8c\xac\xe5\xaa\xe5\x02\x04\x00\x00\xff\xff\x54\x7c\x7d\xf6\xda\x00\x00\x00")

func _lgraphqlSshendpointsbyprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlSshendpointsbyprojectGraphql,
		"_lgraphql/sshEndpointsByProject.graphql",
	)
}

func _lgraphqlSshendpointsbyprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlSshendpointsbyprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/sshEndpointsByProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUpdatedeploytargetGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\x41\x6e\xea\x30\x10\x86\xf7\x73\x8a\x79\x12\x0b\x90\x38\x41\xb6\xaf\x8b\xd2\x05\xa0\xd2\x1e\xc0\xc5\x43\xb0\x9a\x78\x22\x67\x52\x09\x55\xdc\xbd\x0a\x24\xb1\x27\xce\x2a\xf1\xff\x8f\xa5\xf1\xf7\xd5\x9d\x18\x71\xec\x71\x0d\x88\x2b\x67\x0b\xdc\x79\xf9\xd7\xff\x7b\x53\x53\x81\x27\x09\xce\x97\xdb\x3e\x38\xb3\x6f\xb9\xa2\xcf\x50\xa9\x58\xf8\x9b\xbc\x4a\x02\x77\x42\xe1\x68\x44\x28\xe8\xa6\x6d\xaf\xaf\xdc\xca\x3c\x3b\x72\xd0\x59\xcd\xde\x09\xf7\xc7\xff\xec\x2f\xae\x2c\xf0\xed\x74\xd8\x3f\xaa\xaf\xce\x55\x76\x57\x9b\x52\x2f\x77\x09\x8e\xbc\xad\x6e\xfb\x6c\xeb\x8a\x3b\x7b\x0c\xfc\xe3\x2c\x85\xbc\x79\xa7\xd2\xf1\xb4\xe4\x06\x7f\x01\x11\xb1\x6b\xac\x11\x7a\xa1\xa6\xe2\xdb\x87\x09\x25\x49\x31\x64\x87\x86\x7c\x7b\x75\x17\x59\x3b\xdf\x74\x52\x0c\x17\x10\x7b\x74\x2b\x67\x87\x53\x63\xe4\x7c\x8d\x25\xe2\x93\xe6\x03\xea\x76\x0a\x53\xa2\x09\xde\x38\x30\xb0\x7d\x32\x9e\xd2\x19\x5f\xcd\x3b\x5e\x9e\x60\x8f\xd8\x55\xf5\x64\x3e\xd2\x8f\x55\xca\x37\x81\x1d\x07\x72\x37\x99\xae\x38\xac\xb5\x28\x4b\x09\x06\xad\x48\x2b\x9b\x8d\x8d\xbe\x52\x7b\xc3\xc4\xfd\xf1\xbd\x6f\x12\x25\x10\xe1\x0f\xbf\xe7\x40\x46\x68\x2c\x52\xae\x91\x3f\x28\x7e\xa0\x90\xc1\x9c\x12\x2c\xbc\x01\xb2\x85\x21\xc7\x01\x0b\x32\x61\x19\x31\x8c\xcf\xbb\xc3\x5f\x00\x00\x00\xff\xff\x7f\xe0\xea\x58\xb0\x03\x00\x00")

func _lgraphqlUpdatedeploytargetGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUpdatedeploytargetGraphql,
		"_lgraphql/updateDeployTarget.graphql",
	)
}

func _lgraphqlUpdatedeploytargetGraphql() (*asset, error) {
	bytes, err := _lgraphqlUpdatedeploytargetGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/updateDeployTarget.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUpdatedeploytargetconfigGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\x41\x4e\xf4\x30\x0c\x85\xf7\x39\x45\x7e\x69\x16\x33\x57\xe8\xf6\x67\xc3\x06\x55\xc0\x05\x42\xe3\x49\x8d\x32\x4e\xf0\x38\xa0\x11\xea\xdd\x51\x1b\x28\x71\x19\xf0\xaa\x79\xef\xd5\x7a\xfe\x4e\x45\x9c\x60\x22\xbb\x37\xd6\x5a\xbb\x43\xdf\xd9\x5b\x92\x7f\xf5\xf5\x06\x18\x46\x59\x94\x2a\x3c\xb1\xa3\x61\x84\x73\x67\x1f\x84\x91\x42\x55\x73\x89\x91\xe1\xa5\xc0\x59\x36\x8e\x87\x1c\xd3\xe5\xd1\x71\x80\x76\x4d\x2b\xf7\x9c\x9e\x61\x90\xde\x89\x00\xd3\xfa\xfb\xe1\x7d\x89\x96\xec\x9d\xc0\x4d\x93\xff\x9f\xe8\x88\x61\x8f\x94\x8b\x74\x35\x34\xcf\x5c\x7c\x87\x7e\x7d\x67\x27\xc3\xd8\xf8\xf3\xe8\x36\xaa\x85\xca\x7d\x9d\xfd\x79\xbf\xf2\xbe\x09\xac\x30\x94\xaf\x59\x28\x34\xbf\x76\xd9\x22\xf8\x83\xcf\xba\x63\x5a\xbe\xa6\x43\x4b\xc0\x5c\x5b\xae\x11\x34\xa1\x79\xc8\x9d\x40\x09\x47\x46\x20\x1f\x2f\x77\x5b\x63\x88\xa9\xf8\x9e\xd3\x2b\x7a\xe0\x9f\xce\x3d\x04\x4c\xdb\x76\x2d\x2f\x73\x0d\x90\xd1\xc4\xeb\x49\x66\xfa\x08\x00\x00\xff\xff\x82\xe7\x64\x80\x95\x02\x00\x00")

func _lgraphqlUpdatedeploytargetconfigGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUpdatedeploytargetconfigGraphql,
		"_lgraphql/updateDeployTargetConfig.graphql",
	)
}

func _lgraphqlUpdatedeploytargetconfigGraphql() (*asset, error) {
	bytes, err := _lgraphqlUpdatedeploytargetconfigGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/updateDeployTargetConfig.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlVariablesAddorupdateenvvariablebynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\xcc\x2b\x28\x2d\xb1\x52\x70\xcd\x2b\x0b\x4b\x2c\xca\x4c\x4c\xca\x49\x75\xaa\xf4\x4b\xcc\x4d\xf5\x04\x89\x2b\x72\x69\x2a\x54\x73\x29\x28\x28\x28\x24\xa6\xa4\xf8\x17\x85\x16\xa4\x24\x96\xa4\x62\x28\xd5\x80\x9a\x01\x31\x0b\xa6\x03\x04\x32\x53\xe0\xcc\xbc\xc4\xdc\x54\x38\xa7\x2c\x31\xa7\x14\xc1\x2b\x4e\xce\x2f\x80\xf0\x6a\xb9\x6a\xb9\x00\x01\x00\x00\xff\xff\x2e\x2c\x3b\xed\x9f\x00\x00\x00")

func _lgraphqlVariablesAddorupdateenvvariablebynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlVariablesAddorupdateenvvariablebynameGraphql,
		"_lgraphql/variables/addOrUpdateEnvVariableByName.graphql",
	)
}

func _lgraphqlVariablesAddorupdateenvvariablebynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlVariablesAddorupdateenvvariablebynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/variables/addOrUpdateEnvVariableByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlVariablesDeleteenvvariablebynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\xcc\x2b\x28\x2d\xb1\x52\x70\x49\xcd\x49\x2d\x49\x75\xcd\x2b\x0b\x4b\x2c\xca\x4c\x4c\xca\x49\x75\xaa\xf4\x4b\xcc\x4d\xf5\x04\xc9\x2a\x72\x69\x2a\x54\x73\x29\x28\x28\x28\xa4\x60\x57\xa5\x01\x35\x04\x62\x98\x26\x57\x2d\x17\x20\x00\x00\xff\xff\x39\x99\x9e\x59\x64\x00\x00\x00")

func _lgraphqlVariablesDeleteenvvariablebynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlVariablesDeleteenvvariablebynameGraphql,
		"_lgraphql/variables/deleteEnvVariableByName.graphql",
	)
}

func _lgraphqlVariablesDeleteenvvariablebynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlVariablesDeleteenvvariablebynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/variables/deleteEnvVariableByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlVariablesGetenvvariablesbyprojectenvironmentnameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\xe0\x52\x50\x50\xc9\xcc\x2b\x28\x2d\xb1\x52\x70\xcd\x2b\x0b\x4b\x2c\xca\x4c\x4c\xca\x49\x75\xaa\x0c\x28\xca\xcf\x4a\x4d\x2e\x71\xcd\x2b\xcb\x2c\xca\xcf\xcb\x4d\xcd\x2b\xf1\x4b\xcc\x4d\xf5\x04\xa9\x54\xe4\xd2\xac\xe6\x52\x50\x50\x50\x48\x4f\x2d\x41\xd2\x53\x8c\x4b\x93\x06\xd4\x7c\x88\x3d\x9a\x0a\x10\xcd\x20\x90\x99\x02\x67\xe6\x25\xe6\xa6\xc2\x39\x65\x89\x39\xa5\x08\x5e\x71\x72\x7e\x01\x84\x57\xcb\x55\xcb\x05\x08\x00\x00\xff\xff\x72\x77\xad\xf4\xb8\x00\x00\x00")

func _lgraphqlVariablesGetenvvariablesbyprojectenvironmentnameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlVariablesGetenvvariablesbyprojectenvironmentnameGraphql,
		"_lgraphql/variables/getEnvVariablesByProjectEnvironmentName.graphql",
	)
}

func _lgraphqlVariablesGetenvvariablesbyprojectenvironmentnameGraphql() (*asset, error) {
	bytes, err := _lgraphqlVariablesGetenvvariablesbyprojectenvironmentnameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/variables/getEnvVariablesByProjectEnvironmentName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsAddnotificationemailGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x89\xa4\xe6\x26\x66\xe6\x38\xa6\xa4\x14\xa5\x16\x17\xc3\x65\x34\x15\xaa\xb9\x14\x14\x14\x14\x12\x53\x52\xfc\xf2\x4b\x32\xd3\x32\x93\xc1\x66\xb8\x82\xd4\x6a\x64\xe6\x15\x94\x96\x58\x41\x55\x28\x28\x40\x8c\x04\x9b\xac\x03\x15\x42\x35\x13\xc5\x0a\xb0\x8a\x5a\x4d\xb8\xee\xcc\x14\x24\x63\xb0\x68\x87\xa8\xe7\x02\x61\x40\x00\x00\x00\xff\xff\x7d\x85\xc8\xa0\xca\x00\x00\x00")

func _lgraphqlNotificationsAddnotificationemailGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsAddnotificationemailGraphql,
		"_lgraphql/notifications/addNotificationEmail.graphql",
	)
}

func _lgraphqlNotificationsAddnotificationemailGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsAddnotificationemailGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/addNotificationEmail.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsAddnotificationmicrosoftteamsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x89\x94\xa7\x26\x65\xe4\xe7\x67\xc3\x05\x35\x15\xaa\xb9\x14\x14\x14\x14\x12\x53\x52\xfc\xf2\x4b\x32\xd3\x32\x93\xc1\xda\x7d\x33\x93\x8b\xf2\x8b\xf3\xd3\x4a\x42\x52\x13\x73\x8b\x35\x32\xf3\x0a\x4a\x4b\xac\xa0\x4a\x15\x14\x20\xc6\x82\x4d\xd7\x81\x0a\xc1\xcd\x85\xd9\x00\x16\xaf\xd5\x84\xeb\xc9\x4c\x41\xd2\x8c\xaa\x09\xa2\x94\x0b\x84\x01\x01\x00\x00\xff\xff\x66\x70\xe5\xad\xbf\x00\x00\x00")

func _lgraphqlNotificationsAddnotificationmicrosoftteamsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsAddnotificationmicrosoftteamsGraphql,
		"_lgraphql/notifications/addNotificationMicrosoftTeams.graphql",
	)
}

func _lgraphqlNotificationsAddnotificationmicrosoftteamsGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsAddnotificationmicrosoftteamsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/addNotificationMicrosoftTeams.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsAddnotificationrocketchatGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8d\x31\xce\xc2\x30\x0c\x46\xf7\x9c\xe2\xfb\xa5\x7f\x68\xa5\x9e\xa0\x2b\x3b\x03\x9c\xc0\x24\x81\x58\xa5\x36\x42\xae\x18\x10\x77\x47\x4d\x9b\x88\x0c\x5e\x9e\xed\xf7\xe6\xc5\xc8\x58\x05\x9d\x03\xfe\x85\xe6\x38\xe2\x6c\x4f\x96\xdb\xdf\xb0\x12\x9f\x48\x24\xde\x5b\xf8\x8a\x97\xa4\x3a\x55\xd8\xe3\xed\x00\x80\x42\x38\xaa\xf1\x95\x7d\x76\x9e\xd4\x4f\xd1\x0e\x89\xac\x63\x79\x2c\x36\xee\x67\xc0\xd6\xc9\xb9\x61\x47\x35\x54\x92\x65\x51\x63\x25\x9b\xf9\xa7\xaf\x32\x0e\x3f\xd6\xd6\xd6\x2a\xb6\x47\xb7\xce\x37\x00\x00\xff\xff\x49\x05\x99\xf8\xf7\x00\x00\x00")

func _lgraphqlNotificationsAddnotificationrocketchatGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsAddnotificationrocketchatGraphql,
		"_lgraphql/notifications/addNotificationRocketChat.graphql",
	)
}

func _lgraphqlNotificationsAddnotificationrocketchatGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsAddnotificationrocketchatGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/addNotificationRocketChat.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsAddnotificationslackGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8d\x31\x0e\x83\x30\x0c\x45\xf7\x9c\xe2\x57\xea\x00\x12\x27\xe0\x10\x5d\x38\x81\x9b\xa4\xc5\x02\x9c\xaa\x32\xea\x50\xf5\xee\x15\x81\x58\x64\xf0\xf2\x6c\xbf\xb7\xac\x4a\xca\x49\xd0\x38\xe0\x2a\xb4\xc4\x1e\x83\xbe\x59\x9e\x97\x6e\x23\x7e\x24\x91\x38\xd7\xf0\x13\xef\x63\x4a\x93\xc1\x16\x5f\x07\x00\x14\xc2\x2d\x29\x3f\xd8\x67\xe7\x30\x93\x9f\x1a\x96\xd7\xaa\xfd\x71\x01\xec\x89\x5c\xea\x0e\x64\x8d\x52\x2b\x0b\xeb\x94\x62\xe6\xbf\xd6\x64\x1c\x4e\xd6\xda\x56\x2b\xf6\x47\xb7\xcd\x3f\x00\x00\xff\xff\x50\xa8\xad\xd3\xf2\x00\x00\x00")

func _lgraphqlNotificationsAddnotificationslackGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsAddnotificationslackGraphql,
		"_lgraphql/notifications/addNotificationSlack.graphql",
	)
}

func _lgraphqlNotificationsAddnotificationslackGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsAddnotificationslackGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/addNotificationSlack.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsAddnotificationtoprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x4d\x0a\x02\x31\x0c\x85\xf7\x39\xc5\x13\x5c\x38\x57\xe8\x21\x06\x41\x2f\x50\xa6\x55\x22\x4c\x5a\x86\xcc\x42\x86\xde\x5d\xda\xfa\x83\xc6\x55\xcb\xcb\x97\xbc\x6f\x5e\xd5\x2b\x27\xc1\x81\x80\x7d\x5e\xd2\x2d\x4e\xea\x70\xd2\x85\xe5\xba\xab\x99\x24\xe5\x0b\x4f\x8d\x3a\xdf\x73\x74\x18\x7f\x12\x83\x8d\x7e\x8e\x9f\x1b\xc3\x46\x80\x0f\xe1\x6b\x2d\x1d\x7b\x55\xad\x05\x58\xf2\xaa\x6e\x6b\x7f\xe0\x6d\xf1\xf2\x79\xe6\xd6\xc4\xc8\xfd\x21\xbb\x8c\xf1\x6b\x64\x21\x60\xe8\xb5\x1c\xda\x23\x7d\x54\xa8\xd0\x23\x00\x00\xff\xff\x6a\x63\xe5\x98\x1b\x01\x00\x00")

func _lgraphqlNotificationsAddnotificationtoprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsAddnotificationtoprojectGraphql,
		"_lgraphql/notifications/addNotificationToProject.graphql",
	)
}

func _lgraphqlNotificationsAddnotificationtoprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsAddnotificationtoprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/addNotificationToProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsAddnotificationwebhookGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x89\x94\xa7\x26\x65\xe4\xe7\x67\xc3\x05\x35\x15\xaa\xb9\x14\x14\x14\x14\x12\x53\x52\xfc\xf2\x4b\x32\xd3\x32\x93\xc1\xda\xc3\x21\xca\x34\x32\xf3\x0a\x4a\x4b\xac\xa0\x6a\x14\x14\x20\xe6\x81\x8d\xd5\x81\x0a\xc1\x0d\x84\x19\x0d\x16\xaf\xd5\x84\xeb\xc9\x4c\x41\xd2\x8c\xaa\x09\xa2\x94\x0b\x84\x01\x01\x00\x00\xff\xff\x49\x84\xcd\x6a\xb8\x00\x00\x00")

func _lgraphqlNotificationsAddnotificationwebhookGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsAddnotificationwebhookGraphql,
		"_lgraphql/notifications/addNotificationWebhook.graphql",
	)
}

func _lgraphqlNotificationsAddnotificationwebhookGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsAddnotificationwebhookGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/addNotificationWebhook.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsDeletenotificationemailGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x49\xcd\x49\x2d\x49\xf5\xcb\x2f\xc9\x4c\xcb\x4c\x06\xab\xb4\xc2\x22\xe6\x9a\x9b\x98\x99\xa3\x91\x99\x57\x50\x5a\x62\x55\x0d\x31\x07\x6c\x5c\xad\x26\x57\x2d\x20\x00\x00\xff\xff\xac\x83\xb7\xaa\x62\x00\x00\x00")

func _lgraphqlNotificationsDeletenotificationemailGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsDeletenotificationemailGraphql,
		"_lgraphql/notifications/deleteNotificationEmail.graphql",
	)
}

func _lgraphqlNotificationsDeletenotificationemailGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsDeletenotificationemailGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/deleteNotificationEmail.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsDeletenotificationmicrosoftteamsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x49\xcd\x49\x2d\x49\xf5\xcb\x2f\xc9\x4c\xcb\x4c\x06\xab\xb4\xc2\x22\xe6\x9b\x99\x5c\x94\x5f\x9c\x9f\x56\x12\x92\x9a\x98\x5b\xac\x91\x99\x57\x50\x5a\x62\x55\x0d\x31\x10\x6c\x6e\xad\x26\x57\x2d\x20\x00\x00\xff\xff\x54\xf6\x18\xff\x6b\x00\x00\x00")

func _lgraphqlNotificationsDeletenotificationmicrosoftteamsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsDeletenotificationmicrosoftteamsGraphql,
		"_lgraphql/notifications/deleteNotificationMicrosoftTeams.graphql",
	)
}

func _lgraphqlNotificationsDeletenotificationmicrosoftteamsGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsDeletenotificationmicrosoftteamsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/deleteNotificationMicrosoftTeams.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsDeletenotificationrocketchatGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x49\xcd\x49\x2d\x49\xf5\xcb\x2f\xc9\x4c\xcb\x4c\x06\xab\xb4\xc2\x22\x16\x94\x9f\x9c\x9d\x5a\xe2\x9c\x91\x58\xa2\x91\x99\x57\x50\x5a\x62\x55\x0d\x31\x0c\x6c\x66\xad\x26\x57\x2d\x20\x00\x00\xff\xff\x36\x65\x00\x24\x67\x00\x00\x00")

func _lgraphqlNotificationsDeletenotificationrocketchatGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsDeletenotificationrocketchatGraphql,
		"_lgraphql/notifications/deleteNotificationRocketChat.graphql",
	)
}

func _lgraphqlNotificationsDeletenotificationrocketchatGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsDeletenotificationrocketchatGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/deleteNotificationRocketChat.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsDeletenotificationslackGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x49\xcd\x49\x2d\x49\xf5\xcb\x2f\xc9\x4c\xcb\x4c\x06\xab\xb4\xc2\x22\x16\x9c\x93\x98\x9c\xad\x91\x99\x57\x50\x5a\x62\x55\x0d\x31\x07\x6c\x5c\xad\x26\x57\x2d\x20\x00\x00\xff\xff\xee\xc5\xad\x57\x62\x00\x00\x00")

func _lgraphqlNotificationsDeletenotificationslackGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsDeletenotificationslackGraphql,
		"_lgraphql/notifications/deleteNotificationSlack.graphql",
	)
}

func _lgraphqlNotificationsDeletenotificationslackGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsDeletenotificationslackGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/deleteNotificationSlack.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsDeletenotificationwebhookGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x49\xcd\x49\x2d\x49\xf5\xcb\x2f\xc9\x4c\xcb\x4c\x06\xab\xb4\xc2\x22\x16\x9e\x9a\x94\x91\x9f\x9f\xad\x91\x99\x57\x50\x5a\x62\x55\x0d\x31\x09\x6c\x60\xad\x26\x57\x2d\x20\x00\x00\xff\xff\x8a\x89\x7f\x55\x64\x00\x00\x00")

func _lgraphqlNotificationsDeletenotificationwebhookGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsDeletenotificationwebhookGraphql,
		"_lgraphql/notifications/deleteNotificationWebhook.graphql",
	)
}

func _lgraphqlNotificationsDeletenotificationwebhookGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsDeletenotificationwebhookGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/deleteNotificationWebhook.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsListallnotificationemailGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x50\x50\x48\xcc\xc9\x09\x28\xca\xcf\x4a\x4d\x2e\x29\x86\x8a\x80\x40\x5e\x62\x6e\x2a\x9c\x93\x99\x82\x10\xcf\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x2b\xd6\x28\xa9\x2c\x48\xb5\x52\x70\xf5\x75\xf4\xf4\xd1\x44\xd2\x0c\x02\x7a\x7a\x7a\x0a\xf9\x79\x0a\x7e\x48\xea\x5d\x73\x13\x33\x73\xd0\x94\x81\x40\x7c\x3c\xc8\x1c\x14\x1b\x61\x20\x15\xa4\xc5\x31\x25\xa5\x28\xb5\xb8\x18\x43\x12\x43\x47\x2d\x17\x2a\xab\x96\xab\x16\x10\x00\x00\xff\xff\x24\xd5\x22\x2e\xea\x00\x00\x00")

func _lgraphqlNotificationsListallnotificationemailGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsListallnotificationemailGraphql,
		"_lgraphql/notifications/listAllNotificationEmail.graphql",
	)
}

func _lgraphqlNotificationsListallnotificationemailGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsListallnotificationemailGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/listAllNotificationEmail.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsListallnotificationmicrosoftteamsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x50\x50\x48\xcc\xc9\x09\x28\xca\xcf\x4a\x4d\x2e\x29\x86\x8a\x80\x40\x5e\x62\x6e\x2a\x9c\x93\x99\x82\x10\xcf\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x2b\xd6\x28\xa9\x2c\x48\xb5\x52\xf0\xf5\x74\x0e\xf2\x0f\xf6\x77\x0b\x09\x71\x75\xf4\x0d\xd6\x44\x32\x05\x04\xf4\xf4\xf4\x14\xf2\xf3\x14\xfc\x90\x34\xfa\x66\x26\x17\xe5\x17\xe7\xa7\x95\x84\xa4\x26\xe6\x16\xa3\xa9\x07\x81\xf8\x78\x90\xc9\x28\x6e\x80\x81\xf2\xd4\xa4\x8c\xfc\xfc\x6c\x0c\x71\x0c\xc5\xb5\x5c\xa8\xac\x5a\xae\x5a\x40\x00\x00\x00\xff\xff\x3c\x1a\x87\x3b\xf7\x00\x00\x00")

func _lgraphqlNotificationsListallnotificationmicrosoftteamsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsListallnotificationmicrosoftteamsGraphql,
		"_lgraphql/notifications/listAllNotificationMicrosoftTeams.graphql",
	)
}

func _lgraphqlNotificationsListallnotificationmicrosoftteamsGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsListallnotificationmicrosoftteamsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/listAllNotificationMicrosoftTeams.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsListallnotificationrocketchatGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x50\x50\x48\xcc\xc9\x09\x28\xca\xcf\x4a\x4d\x2e\x29\x86\x8a\x80\x40\x5e\x62\x6e\x2a\x9c\x93\x99\x82\x10\xcf\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x2b\xd6\x28\xa9\x2c\x48\xb5\x52\x08\xf2\x77\xf6\x76\x0d\x71\xf6\x70\x0c\xd1\x44\x32\x01\x04\xf4\xf4\xf4\x14\xf2\xf3\x14\xfc\x90\x34\x05\xe5\x27\x67\xa7\x96\x38\x67\x24\x96\xa0\xa9\x05\x81\xf8\x78\x90\x89\x28\x76\xc3\x40\x79\x6a\x52\x46\x7e\x7e\x36\x86\x78\x72\x46\x62\x5e\x5e\x6a\x0e\x86\x38\x86\x21\xb5\x5c\xa8\xac\x5a\xae\x5a\x40\x00\x00\x00\xff\xff\x4a\x62\x23\x28\x07\x01\x00\x00")

func _lgraphqlNotificationsListallnotificationrocketchatGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsListallnotificationrocketchatGraphql,
		"_lgraphql/notifications/listAllNotificationRocketChat.graphql",
	)
}

func _lgraphqlNotificationsListallnotificationrocketchatGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsListallnotificationrocketchatGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/listAllNotificationRocketChat.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsListallnotificationslackGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x50\x50\x48\xcc\xc9\x09\x28\xca\xcf\x4a\x4d\x2e\x29\x86\x8a\x80\x40\x5e\x62\x6e\x2a\x9c\x93\x99\x82\x10\xcf\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x2b\xd6\x28\xa9\x2c\x48\xb5\x52\x08\xf6\x71\x74\xf6\xd6\x44\xd2\x0c\x02\x7a\x7a\x7a\x0a\xf9\x79\x0a\x7e\x48\xea\x83\x73\x12\x93\xb3\xd1\x94\x81\x40\x7c\x3c\xc8\x1c\x14\x1b\x61\xa0\x3c\x35\x29\x23\x3f\x3f\x1b\x43\x3c\x39\x23\x31\x2f\x2f\x35\x07\x43\x1c\xc3\x90\x5a\x2e\x54\x56\x2d\x57\x2d\x20\x00\x00\xff\xff\xff\x6b\xa2\xcc\xfd\x00\x00\x00")

func _lgraphqlNotificationsListallnotificationslackGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsListallnotificationslackGraphql,
		"_lgraphql/notifications/listAllNotificationSlack.graphql",
	)
}

func _lgraphqlNotificationsListallnotificationslackGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsListallnotificationslackGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/listAllNotificationSlack.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsListallnotificationwebhookGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x50\x50\x48\xcc\xc9\x09\x28\xca\xcf\x4a\x4d\x2e\x29\x86\x8a\x80\x40\x5e\x62\x6e\x2a\x9c\x93\x99\x82\x10\xcf\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x2b\xd6\x28\xa9\x2c\x48\xb5\x52\x08\x77\x75\xf2\xf0\xf7\xf7\xd6\x44\xd2\x0e\x02\x7a\x7a\x7a\x0a\xf9\x79\x0a\x7e\x48\x3a\xc2\x53\x93\x32\xf2\xf3\xb3\xd1\x14\x82\x40\x7c\x3c\xc8\x2c\x14\x5b\x61\xa0\x1c\xa2\x09\x43\x1c\x43\x71\x2d\x17\x2a\xab\x96\xab\x16\x10\x00\x00\xff\xff\x92\xfc\x32\xab\xe9\x00\x00\x00")

func _lgraphqlNotificationsListallnotificationwebhookGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsListallnotificationwebhookGraphql,
		"_lgraphql/notifications/listAllNotificationWebhook.graphql",
	)
}

func _lgraphqlNotificationsListallnotificationwebhookGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsListallnotificationwebhookGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/listAllNotificationWebhook.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsProjectnotificationemailGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x71\xaa\xf4\x4b\xcc\x4d\xd5\x80\x28\x01\xab\x84\x29\x00\x81\xbc\xfc\x92\xcc\xb4\xcc\xe4\xc4\x92\xcc\xfc\xbc\x62\x8d\x92\xca\x82\x54\x2b\x05\x57\x5f\x47\x4f\x1f\x64\x45\x20\xa0\xa7\xa7\xa7\x90\x9f\xa7\xe0\x87\xa4\xde\x35\x37\x31\x33\x07\x4d\x19\x08\xc4\xc7\x83\xcc\x01\xd9\x84\x21\x95\x0a\xd2\xe2\x98\x92\x52\x94\x5a\x5c\x8c\x21\x89\xa1\xa3\x96\x0b\x95\x55\xcb\x55\x0b\x08\x00\x00\xff\xff\x65\xff\x3e\x1e\xf2\x00\x00\x00")

func _lgraphqlNotificationsProjectnotificationemailGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsProjectnotificationemailGraphql,
		"_lgraphql/notifications/projectNotificationEmail.graphql",
	)
}

func _lgraphqlNotificationsProjectnotificationemailGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsProjectnotificationemailGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/projectNotificationEmail.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsProjectnotificationmicrosoftteamsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x71\xaa\xf4\x4b\xcc\x4d\xd5\x80\x28\x01\xab\x84\x29\x00\x81\xbc\xfc\x92\xcc\xb4\xcc\xe4\xc4\x92\xcc\xfc\xbc\x62\x8d\x92\xca\x82\x54\x2b\x05\x5f\x4f\xe7\x20\xff\x60\x7f\xb7\x90\x10\x57\x47\xdf\x60\x64\xd5\x20\xa0\xa7\xa7\xa7\x90\x9f\xa7\xe0\x87\xa4\xd1\x37\x33\xb9\x28\xbf\x38\x3f\xad\x24\x24\x35\x31\xb7\x18\x4d\x3d\x08\xc4\xc7\x83\x4c\x06\xd9\x8d\x21\x55\x9e\x9a\x94\x91\x9f\x9f\x8d\x21\x8e\xa1\xb8\x96\x0b\x95\x55\xcb\x55\x0b\x08\x00\x00\xff\xff\xbf\xb5\x85\x65\xff\x00\x00\x00")

func _lgraphqlNotificationsProjectnotificationmicrosoftteamsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsProjectnotificationmicrosoftteamsGraphql,
		"_lgraphql/notifications/projectNotificationMicrosoftTeams.graphql",
	)
}

func _lgraphqlNotificationsProjectnotificationmicrosoftteamsGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsProjectnotificationmicrosoftteamsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/projectNotificationMicrosoftTeams.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsProjectnotificationrocketchatGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\xc1\x0a\xc2\x30\x0c\x86\xef\x7b\x8a\x08\x1e\xe6\x65\x0f\xb0\x9b\x0e\x41\x10\x26\xcc\xdd\x47\x2d\xd1\xd6\xb9\x64\xd6\x88\x14\xe9\xbb\x4b\x15\x61\xda\xef\x14\xfe\x7c\xf9\x73\xbd\xa3\xf3\x90\xcf\x49\x0d\x58\xc2\x5e\x9c\xa5\xd3\x6c\x01\xcf\x0c\x00\x60\x74\x7c\x46\x2d\x2b\x5f\xab\x01\xf3\x8f\xf2\x36\xbf\x42\x84\x58\xec\xd1\x6a\x25\x96\xe9\x96\x8b\x1f\xb1\x84\x66\x57\x6d\xd7\x6d\xb5\x59\xb6\x53\x33\x52\x14\x05\x30\x41\x3d\x39\x6a\x58\xf7\x28\x95\x51\xf2\xe7\x46\xba\x2e\x36\xc6\x9f\xc9\xea\x81\x07\xc3\xdc\x27\xb9\x36\x8a\x08\x2f\x49\x9e\x94\x84\xec\x77\x0a\x59\x78\x05\x00\x00\xff\xff\x98\xc9\xde\xdd\x0f\x01\x00\x00")

func _lgraphqlNotificationsProjectnotificationrocketchatGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsProjectnotificationrocketchatGraphql,
		"_lgraphql/notifications/projectNotificationRocketChat.graphql",
	)
}

func _lgraphqlNotificationsProjectnotificationrocketchatGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsProjectnotificationrocketchatGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/projectNotificationRocketChat.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsProjectnotificationslackGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x71\xaa\xf4\x4b\xcc\x4d\xd5\x80\x28\x01\xab\x84\x29\x00\x81\xbc\xfc\x92\xcc\xb4\xcc\xe4\xc4\x92\xcc\xfc\xbc\x62\x8d\x92\xca\x02\x90\x39\x3e\x8e\xce\xde\xc8\x8a\x40\x40\x4f\x4f\x4f\x21\x3f\x4f\xc1\x0f\x49\x7d\x70\x4e\x62\x72\x36\x9a\x32\x10\x88\x8f\x07\x99\x03\xb2\x09\x43\xaa\x3c\x35\x29\x23\x3f\x3f\x1b\x43\x3c\x39\x23\x31\x2f\x2f\x35\x07\x43\x1c\xc3\x90\x5a\x2e\x54\x56\x2d\x57\x2d\x20\x00\x00\xff\xff\xff\xce\xda\x6a\x05\x01\x00\x00")

func _lgraphqlNotificationsProjectnotificationslackGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsProjectnotificationslackGraphql,
		"_lgraphql/notifications/projectNotificationSlack.graphql",
	)
}

func _lgraphqlNotificationsProjectnotificationslackGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsProjectnotificationslackGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/projectNotificationSlack.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsProjectnotificationwebhookGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x71\xaa\xf4\x4b\xcc\x4d\xd5\x80\x28\x01\xab\x84\x29\x00\x81\xbc\xfc\x92\xcc\xb4\xcc\xe4\xc4\x92\xcc\xfc\xbc\x62\x8d\x92\xca\x82\x54\x2b\x85\x70\x57\x27\x0f\x7f\x7f\x6f\x64\x65\x20\xa0\xa7\xa7\xa7\x90\x9f\xa7\xe0\x87\xa4\x23\x3c\x35\x29\x23\x3f\x3f\x1b\x4d\x21\x08\xc4\xc7\x83\xcc\x02\xd9\x86\x21\x55\x0e\xd1\x84\x21\x8e\xa1\xb8\x96\x0b\x95\x55\xcb\x55\x0b\x08\x00\x00\xff\xff\x7d\x54\xe7\x02\xf1\x00\x00\x00")

func _lgraphqlNotificationsProjectnotificationwebhookGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsProjectnotificationwebhookGraphql,
		"_lgraphql/notifications/projectNotificationWebhook.graphql",
	)
}

func _lgraphqlNotificationsProjectnotificationwebhookGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsProjectnotificationwebhookGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/projectNotificationWebhook.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsRemovenotificationfromprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x4d\x0a\x02\x31\x0c\x85\xf7\x39\xc5\x13\x5c\x38\x57\xe8\x01\x5c\x0e\x82\x5e\x60\x18\xa3\x44\x68\x53\x4a\x46\x90\x61\xee\x2e\x6d\xfd\xc3\xce\x2e\x79\x7c\xc9\xfb\xfc\x64\x83\x89\x06\xec\x08\xd8\xc6\xa4\x37\x1e\xcd\xe1\x68\x49\xc2\x75\x93\xb3\xa0\x26\x17\x19\x0b\x75\x7a\x44\x76\xe8\xff\x92\x06\xeb\x07\xcf\xdf\x1f\xdd\x4c\x40\x62\xaf\x77\xfe\xbd\xdc\x27\xf5\x87\xda\x97\xbb\x01\x09\x71\x32\x37\x97\x19\xf8\xa8\xbc\xa5\x5e\x79\xab\xd3\x18\xae\x90\xd5\xa8\x91\x2c\xe4\x42\x40\x57\x6b\xe5\x4c\x79\x5f\xe8\x19\x00\x00\xff\xff\x17\x0d\x7b\xbb\x17\x01\x00\x00")

func _lgraphqlNotificationsRemovenotificationfromprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsRemovenotificationfromprojectGraphql,
		"_lgraphql/notifications/removeNotificationFromProject.graphql",
	)
}

func _lgraphqlNotificationsRemovenotificationfromprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsRemovenotificationfromprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/removeNotificationFromProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsUpdatenotificationemailGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x3d\x0a\x02\x31\x10\x85\xeb\xcc\x29\x66\xc1\xc2\xbd\x42\x3a\x0b\x0b\x1b\x11\xc4\x03\x0c\x9b\xa8\x03\x26\x1b\x76\x27\x55\xc8\xdd\x65\x36\x12\x14\x9c\xee\xfd\xe4\x7d\x09\x59\x48\x78\x8e\xb8\x07\x44\xc4\x5d\xa4\xe0\x2d\x5e\x65\xe1\xf8\x18\x9a\x95\x48\xa6\xa7\xc5\x5b\x72\x24\xfe\x3c\x0b\xdf\x79\xda\xde\x1c\x03\xf1\xeb\xa2\xe9\x29\xa6\x2c\x03\x8c\x58\xc0\xe4\xff\xbd\xb6\xaf\xc7\x5a\xb6\xa5\x6b\xbd\x86\xdd\xe8\x3f\xfe\x87\xdd\xfe\xd0\x93\x0a\x66\x2c\x60\x0c\x3b\xf8\x1e\xe8\xc2\x2b\xf0\xe0\xdc\xe2\xd7\x15\x4c\x85\xfa\x0e\x00\x00\xff\xff\x53\xd5\xfd\x0c\xe6\x00\x00\x00")

func _lgraphqlNotificationsUpdatenotificationemailGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsUpdatenotificationemailGraphql,
		"_lgraphql/notifications/updateNotificationEmail.graphql",
	)
}

func _lgraphqlNotificationsUpdatenotificationemailGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsUpdatenotificationemailGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/updateNotificationEmail.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsUpdatenotificationmicrosoftteamsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\x31\x0e\xc2\x30\x0c\x45\xe7\xf8\x14\xae\xc4\x40\xaf\xd0\x1b\x30\x80\x90\x80\x03\x84\x34\xa5\x16\x4a\x1c\xb5\x8e\x18\xa2\xdc\x1d\xa5\x41\x11\x4c\xfc\xed\x7f\xfb\xdb\xcf\x45\xd1\x42\xec\x71\x0f\x88\x88\x3b\xaf\x9d\x1d\xf0\x22\x0b\xf9\x47\x57\xa3\xa0\xc5\xcc\x03\xde\xc2\xa8\xc5\x9e\x58\x68\x22\xb3\x75\x8e\x64\x16\x5e\x79\x92\xab\xd5\x6e\x3d\x97\xb5\x83\x0f\x51\x3a\xe8\x31\x81\x8a\x7f\x0a\xf5\x63\x11\x95\xd6\x90\x9a\x2f\xaa\x20\x1b\xcf\x4f\xfe\xa1\xa9\x54\x6d\x92\x41\xf5\x09\x94\xa2\x11\xbe\x0f\x34\xf3\xb2\xf7\x99\xf9\x09\x2a\x43\x7e\x07\x00\x00\xff\xff\x71\xce\xd3\xe8\xf3\x00\x00\x00")

func _lgraphqlNotificationsUpdatenotificationmicrosoftteamsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsUpdatenotificationmicrosoftteamsGraphql,
		"_lgraphql/notifications/updateNotificationMicrosoftTeams.graphql",
	)
}

func _lgraphqlNotificationsUpdatenotificationmicrosoftteamsGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsUpdatenotificationmicrosoftteamsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/updateNotificationMicrosoftTeams.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsUpdatenotificationrocketchatGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xb1\x0e\xc2\x30\x0c\x44\xe7\xf8\x2b\x5c\x89\x81\xfe\x42\x57\x26\x16\x84\x40\x7c\x80\x49\x03\xb1\x4a\x9d\x0a\x39\x62\x88\xf2\xef\x28\x0d\x8a\x60\xe1\xb6\xbb\xdc\xc5\x6f\x8e\x4a\xca\x41\x70\x0b\x88\x88\x1b\xa1\xd9\x0d\x78\xd6\x27\xcb\xbd\xab\xd1\x42\x6a\xfd\x80\x97\x65\x24\x75\x87\xa0\x7c\x63\xbb\x6e\x4e\xc1\x4e\x4e\x77\x9e\xf4\x58\x2a\x7b\x59\xa2\x76\xd0\x63\x02\x13\xff\x94\xeb\xa5\x22\x2e\x8b\x21\x35\x5f\x54\x01\x56\x8e\x9f\xfc\x43\x51\x69\xda\x4b\x06\xd3\x27\x30\x86\x47\xf8\xfe\xa0\x19\xeb\x49\xc4\x3d\x9a\x7f\xb9\xab\x0f\x61\x02\x93\x21\xbf\x03\x00\x00\xff\xff\x5d\x2a\x5a\x5e\xfb\x00\x00\x00")

func _lgraphqlNotificationsUpdatenotificationrocketchatGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsUpdatenotificationrocketchatGraphql,
		"_lgraphql/notifications/updateNotificationRocketChat.graphql",
	)
}

func _lgraphqlNotificationsUpdatenotificationrocketchatGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsUpdatenotificationrocketchatGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/updateNotificationRocketChat.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsUpdatenotificationslackGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x4e\x4b\x0a\xc2\x30\x10\x5d\x67\x4e\x31\x05\x17\xf6\x0a\xbd\x81\x1b\x11\x8a\x07\x18\xd3\x68\x87\xb6\x93\x20\x13\x5c\x84\xdc\x5d\xd2\x48\x50\xf0\xed\xde\x67\xde\x9b\x2d\x2a\x29\x7b\xc1\x23\x20\x22\x1e\x84\x36\x37\xe0\xa8\x4f\x96\x47\x57\xa5\x40\x6a\xe7\x01\xaf\x61\x22\x75\x67\xaf\x7c\x67\xbb\xdf\x8c\x2b\xd9\xe5\x52\xdc\x93\x84\xa8\x1d\xf4\x98\xc0\xc4\xff\xb9\xda\x5f\xc0\x25\x3c\xa4\xc6\x0b\xea\xec\xbe\xfe\xa3\x7f\xb6\xeb\x0f\xcd\xc9\x60\xfa\x04\xc6\xf0\x04\xdf\x05\x8d\xd8\x99\x44\xdc\xda\xf8\xcb\xdd\x66\xef\x17\x30\x19\xf2\x3b\x00\x00\xff\xff\xd4\x02\x8b\x34\xf1\x00\x00\x00")

func _lgraphqlNotificationsUpdatenotificationslackGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsUpdatenotificationslackGraphql,
		"_lgraphql/notifications/updateNotificationSlack.graphql",
	)
}

func _lgraphqlNotificationsUpdatenotificationslackGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsUpdatenotificationslackGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/updateNotificationSlack.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsUpdatenotificationwebhookGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\x84\x08\x15\x24\x96\x24\x67\x58\x29\x84\x16\xa4\x24\x96\xa4\xfa\xe5\x97\x64\xa6\x65\x26\x83\xf5\x84\xa7\x26\x65\xe4\xe7\x67\x07\x80\xe4\x3d\xf3\x0a\x4a\x4b\x14\xb9\x34\x15\xaa\xb9\x38\x4b\x71\xa9\x84\xd8\x01\x02\x99\x20\xe5\x56\xd5\x70\x3e\x08\x40\xac\x06\xbb\x00\x45\x1c\x6a\x3f\xc4\x1d\x70\x99\x5a\x2e\x4e\xcd\x6a\x2e\x4e\xce\xcc\x14\x2e\x64\x03\xe0\x9c\x72\x88\x95\x5c\x9c\xb5\x5c\xb5\x80\x00\x00\x00\xff\xff\x93\xb3\xe1\xf1\xe5\x00\x00\x00")

func _lgraphqlNotificationsUpdatenotificationwebhookGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsUpdatenotificationwebhookGraphql,
		"_lgraphql/notifications/updateNotificationWebhook.graphql",
	)
}

func _lgraphqlNotificationsUpdatenotificationwebhookGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsUpdatenotificationwebhookGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/updateNotificationWebhook.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"_lgraphql/addDeployTarget.graphql":                                   _lgraphqlAdddeploytargetGraphql,
	"_lgraphql/addEnvVariable.graphql":                                    _lgraphqlAddenvvariableGraphql,
	"_lgraphql/addGroup.graphql":                                          _lgraphqlAddgroupGraphql,
	"_lgraphql/addGroupsToProject.graphql":                                _lgraphqlAddgroupstoprojectGraphql,
	"_lgraphql/addOrUpdateEnvironment.graphql":                            _lgraphqlAddorupdateenvironmentGraphql,
	"_lgraphql/addProject.graphql":                                        _lgraphqlAddprojectGraphql,
	"_lgraphql/addRestore.graphql":                                        _lgraphqlAddrestoreGraphql,
	"_lgraphql/addSshKey.graphql":                                         _lgraphqlAddsshkeyGraphql,
	"_lgraphql/addUser.graphql":                                           _lgraphqlAdduserGraphql,
	"_lgraphql/addUserToGroup.graphql":                                    _lgraphqlAddusertogroupGraphql,
	"_lgraphql/backupsForEnvironmentByName.graphql":                       _lgraphqlBackupsforenvironmentbynameGraphql,
	"_lgraphql/deleteDeployTarget.graphql":                                _lgraphqlDeletedeploytargetGraphql,
	"_lgraphql/deleteDeployTargetConfig.graphql":                          _lgraphqlDeletedeploytargetconfigGraphql,
	"_lgraphql/deployEnvironmentBranch.graphql":                           _lgraphqlDeployenvironmentbranchGraphql,
	"_lgraphql/deployEnvironmentLatest.graphql":                           _lgraphqlDeployenvironmentlatestGraphql,
	"_lgraphql/deployEnvironmentPromote.graphql":                          _lgraphqlDeployenvironmentpromoteGraphql,
	"_lgraphql/deployEnvironmentPullrequest.graphql":                      _lgraphqlDeployenvironmentpullrequestGraphql,
	"_lgraphql/deployTargetConfigsByProjectId.graphql":                    _lgraphqlDeploytargetconfigsbyprojectidGraphql,
	"_lgraphql/environmentByName.graphql":                                 _lgraphqlEnvironmentbynameGraphql,
	"_lgraphql/lagoonSchema.graphql":                                      _lgraphqlLagoonschemaGraphql,
	"_lgraphql/lagoonVersion.graphql":                                     _lgraphqlLagoonversionGraphql,
	"_lgraphql/listDeployTargets.graphql":                                 _lgraphqlListdeploytargetsGraphql,
	"_lgraphql/me.graphql":                                                _lgraphqlMeGraphql,
	"_lgraphql/minimalProjectByName.graphql":                              _lgraphqlMinimalprojectbynameGraphql,
	"_lgraphql/projectByName.graphql":                                     _lgraphqlProjectbynameGraphql,
	"_lgraphql/projectByNameMetadata.graphql":                             _lgraphqlProjectbynamemetadataGraphql,
	"_lgraphql/sshEndpointsByProject.graphql":                             _lgraphqlSshendpointsbyprojectGraphql,
	"_lgraphql/updateDeployTarget.graphql":                                _lgraphqlUpdatedeploytargetGraphql,
	"_lgraphql/updateDeployTargetConfig.graphql":                          _lgraphqlUpdatedeploytargetconfigGraphql,
	"_lgraphql/variables/addOrUpdateEnvVariableByName.graphql":            _lgraphqlVariablesAddorupdateenvvariablebynameGraphql,
	"_lgraphql/variables/deleteEnvVariableByName.graphql":                 _lgraphqlVariablesDeleteenvvariablebynameGraphql,
	"_lgraphql/variables/getEnvVariablesByProjectEnvironmentName.graphql": _lgraphqlVariablesGetenvvariablesbyprojectenvironmentnameGraphql,
	"_lgraphql/notifications/addNotificationEmail.graphql":                _lgraphqlNotificationsAddnotificationemailGraphql,
	"_lgraphql/notifications/addNotificationMicrosoftTeams.graphql":       _lgraphqlNotificationsAddnotificationmicrosoftteamsGraphql,
	"_lgraphql/notifications/addNotificationRocketChat.graphql":           _lgraphqlNotificationsAddnotificationrocketchatGraphql,
	"_lgraphql/notifications/addNotificationSlack.graphql":                _lgraphqlNotificationsAddnotificationslackGraphql,
	"_lgraphql/notifications/addNotificationToProject.graphql":            _lgraphqlNotificationsAddnotificationtoprojectGraphql,
	"_lgraphql/notifications/addNotificationWebhook.graphql":              _lgraphqlNotificationsAddnotificationwebhookGraphql,
	"_lgraphql/notifications/deleteNotificationEmail.graphql":             _lgraphqlNotificationsDeletenotificationemailGraphql,
	"_lgraphql/notifications/deleteNotificationMicrosoftTeams.graphql":    _lgraphqlNotificationsDeletenotificationmicrosoftteamsGraphql,
	"_lgraphql/notifications/deleteNotificationRocketChat.graphql":        _lgraphqlNotificationsDeletenotificationrocketchatGraphql,
	"_lgraphql/notifications/deleteNotificationSlack.graphql":             _lgraphqlNotificationsDeletenotificationslackGraphql,
	"_lgraphql/notifications/deleteNotificationWebhook.graphql":           _lgraphqlNotificationsDeletenotificationwebhookGraphql,
	"_lgraphql/notifications/listAllNotificationEmail.graphql":            _lgraphqlNotificationsListallnotificationemailGraphql,
	"_lgraphql/notifications/listAllNotificationMicrosoftTeams.graphql":   _lgraphqlNotificationsListallnotificationmicrosoftteamsGraphql,
	"_lgraphql/notifications/listAllNotificationRocketChat.graphql":       _lgraphqlNotificationsListallnotificationrocketchatGraphql,
	"_lgraphql/notifications/listAllNotificationSlack.graphql":            _lgraphqlNotificationsListallnotificationslackGraphql,
	"_lgraphql/notifications/listAllNotificationWebhook.graphql":          _lgraphqlNotificationsListallnotificationwebhookGraphql,
	"_lgraphql/notifications/projectNotificationEmail.graphql":            _lgraphqlNotificationsProjectnotificationemailGraphql,
	"_lgraphql/notifications/projectNotificationMicrosoftTeams.graphql":   _lgraphqlNotificationsProjectnotificationmicrosoftteamsGraphql,
	"_lgraphql/notifications/projectNotificationRocketChat.graphql":       _lgraphqlNotificationsProjectnotificationrocketchatGraphql,
	"_lgraphql/notifications/projectNotificationSlack.graphql":            _lgraphqlNotificationsProjectnotificationslackGraphql,
	"_lgraphql/notifications/projectNotificationWebhook.graphql":          _lgraphqlNotificationsProjectnotificationwebhookGraphql,
	"_lgraphql/notifications/removeNotificationFromProject.graphql":       _lgraphqlNotificationsRemovenotificationfromprojectGraphql,
	"_lgraphql/notifications/updateNotificationEmail.graphql":             _lgraphqlNotificationsUpdatenotificationemailGraphql,
	"_lgraphql/notifications/updateNotificationMicrosoftTeams.graphql":    _lgraphqlNotificationsUpdatenotificationmicrosoftteamsGraphql,
	"_lgraphql/notifications/updateNotificationRocketChat.graphql":        _lgraphqlNotificationsUpdatenotificationrocketchatGraphql,
	"_lgraphql/notifications/updateNotificationSlack.graphql":             _lgraphqlNotificationsUpdatenotificationslackGraphql,
	"_lgraphql/notifications/updateNotificationWebhook.graphql":           _lgraphqlNotificationsUpdatenotificationwebhookGraphql,
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
		"addDeployTarget.graphql":                &bintree{_lgraphqlAdddeploytargetGraphql, map[string]*bintree{}},
		"addEnvVariable.graphql":                 &bintree{_lgraphqlAddenvvariableGraphql, map[string]*bintree{}},
		"addGroup.graphql":                       &bintree{_lgraphqlAddgroupGraphql, map[string]*bintree{}},
		"addGroupsToProject.graphql":             &bintree{_lgraphqlAddgroupstoprojectGraphql, map[string]*bintree{}},
		"addOrUpdateEnvironment.graphql":         &bintree{_lgraphqlAddorupdateenvironmentGraphql, map[string]*bintree{}},
		"addProject.graphql":                     &bintree{_lgraphqlAddprojectGraphql, map[string]*bintree{}},
		"addRestore.graphql":                     &bintree{_lgraphqlAddrestoreGraphql, map[string]*bintree{}},
		"addSshKey.graphql":                      &bintree{_lgraphqlAddsshkeyGraphql, map[string]*bintree{}},
		"addUser.graphql":                        &bintree{_lgraphqlAdduserGraphql, map[string]*bintree{}},
		"addUserToGroup.graphql":                 &bintree{_lgraphqlAddusertogroupGraphql, map[string]*bintree{}},
		"backupsForEnvironmentByName.graphql":    &bintree{_lgraphqlBackupsforenvironmentbynameGraphql, map[string]*bintree{}},
		"deleteDeployTarget.graphql":             &bintree{_lgraphqlDeletedeploytargetGraphql, map[string]*bintree{}},
		"deleteDeployTargetConfig.graphql":       &bintree{_lgraphqlDeletedeploytargetconfigGraphql, map[string]*bintree{}},
		"deployEnvironmentBranch.graphql":        &bintree{_lgraphqlDeployenvironmentbranchGraphql, map[string]*bintree{}},
		"deployEnvironmentLatest.graphql":        &bintree{_lgraphqlDeployenvironmentlatestGraphql, map[string]*bintree{}},
		"deployEnvironmentPromote.graphql":       &bintree{_lgraphqlDeployenvironmentpromoteGraphql, map[string]*bintree{}},
		"deployEnvironmentPullrequest.graphql":   &bintree{_lgraphqlDeployenvironmentpullrequestGraphql, map[string]*bintree{}},
		"deployTargetConfigsByProjectId.graphql": &bintree{_lgraphqlDeploytargetconfigsbyprojectidGraphql, map[string]*bintree{}},
		"environmentByName.graphql":              &bintree{_lgraphqlEnvironmentbynameGraphql, map[string]*bintree{}},
		"lagoonSchema.graphql":                   &bintree{_lgraphqlLagoonschemaGraphql, map[string]*bintree{}},
		"lagoonVersion.graphql":                  &bintree{_lgraphqlLagoonversionGraphql, map[string]*bintree{}},
		"listDeployTargets.graphql":              &bintree{_lgraphqlListdeploytargetsGraphql, map[string]*bintree{}},
		"me.graphql":                             &bintree{_lgraphqlMeGraphql, map[string]*bintree{}},
		"minimalProjectByName.graphql":           &bintree{_lgraphqlMinimalprojectbynameGraphql, map[string]*bintree{}},
		"notifications": &bintree{nil, map[string]*bintree{
			"addNotificationEmail.graphql":              &bintree{_lgraphqlNotificationsAddnotificationemailGraphql, map[string]*bintree{}},
			"addNotificationMicrosoftTeams.graphql":     &bintree{_lgraphqlNotificationsAddnotificationmicrosoftteamsGraphql, map[string]*bintree{}},
			"addNotificationRocketChat.graphql":         &bintree{_lgraphqlNotificationsAddnotificationrocketchatGraphql, map[string]*bintree{}},
			"addNotificationSlack.graphql":              &bintree{_lgraphqlNotificationsAddnotificationslackGraphql, map[string]*bintree{}},
			"addNotificationToProject.graphql":          &bintree{_lgraphqlNotificationsAddnotificationtoprojectGraphql, map[string]*bintree{}},
			"addNotificationWebhook.graphql":            &bintree{_lgraphqlNotificationsAddnotificationwebhookGraphql, map[string]*bintree{}},
			"deleteNotificationEmail.graphql":           &bintree{_lgraphqlNotificationsDeletenotificationemailGraphql, map[string]*bintree{}},
			"deleteNotificationMicrosoftTeams.graphql":  &bintree{_lgraphqlNotificationsDeletenotificationmicrosoftteamsGraphql, map[string]*bintree{}},
			"deleteNotificationRocketChat.graphql":      &bintree{_lgraphqlNotificationsDeletenotificationrocketchatGraphql, map[string]*bintree{}},
			"deleteNotificationSlack.graphql":           &bintree{_lgraphqlNotificationsDeletenotificationslackGraphql, map[string]*bintree{}},
			"deleteNotificationWebhook.graphql":         &bintree{_lgraphqlNotificationsDeletenotificationwebhookGraphql, map[string]*bintree{}},
			"listAllNotificationEmail.graphql":          &bintree{_lgraphqlNotificationsListallnotificationemailGraphql, map[string]*bintree{}},
			"listAllNotificationMicrosoftTeams.graphql": &bintree{_lgraphqlNotificationsListallnotificationmicrosoftteamsGraphql, map[string]*bintree{}},
			"listAllNotificationRocketChat.graphql":     &bintree{_lgraphqlNotificationsListallnotificationrocketchatGraphql, map[string]*bintree{}},
			"listAllNotificationSlack.graphql":          &bintree{_lgraphqlNotificationsListallnotificationslackGraphql, map[string]*bintree{}},
			"listAllNotificationWebhook.graphql":        &bintree{_lgraphqlNotificationsListallnotificationwebhookGraphql, map[string]*bintree{}},
			"projectNotificationEmail.graphql":          &bintree{_lgraphqlNotificationsProjectnotificationemailGraphql, map[string]*bintree{}},
			"projectNotificationMicrosoftTeams.graphql": &bintree{_lgraphqlNotificationsProjectnotificationmicrosoftteamsGraphql, map[string]*bintree{}},
			"projectNotificationRocketChat.graphql":     &bintree{_lgraphqlNotificationsProjectnotificationrocketchatGraphql, map[string]*bintree{}},
			"projectNotificationSlack.graphql":          &bintree{_lgraphqlNotificationsProjectnotificationslackGraphql, map[string]*bintree{}},
			"projectNotificationWebhook.graphql":        &bintree{_lgraphqlNotificationsProjectnotificationwebhookGraphql, map[string]*bintree{}},
			"removeNotificationFromProject.graphql":     &bintree{_lgraphqlNotificationsRemovenotificationfromprojectGraphql, map[string]*bintree{}},
			"updateNotificationEmail.graphql":           &bintree{_lgraphqlNotificationsUpdatenotificationemailGraphql, map[string]*bintree{}},
			"updateNotificationMicrosoftTeams.graphql":  &bintree{_lgraphqlNotificationsUpdatenotificationmicrosoftteamsGraphql, map[string]*bintree{}},
			"updateNotificationRocketChat.graphql":      &bintree{_lgraphqlNotificationsUpdatenotificationrocketchatGraphql, map[string]*bintree{}},
			"updateNotificationSlack.graphql":           &bintree{_lgraphqlNotificationsUpdatenotificationslackGraphql, map[string]*bintree{}},
			"updateNotificationWebhook.graphql":         &bintree{_lgraphqlNotificationsUpdatenotificationwebhookGraphql, map[string]*bintree{}},
		}},
		"projectByName.graphql":            &bintree{_lgraphqlProjectbynameGraphql, map[string]*bintree{}},
		"projectByNameMetadata.graphql":    &bintree{_lgraphqlProjectbynamemetadataGraphql, map[string]*bintree{}},
		"sshEndpointsByProject.graphql":    &bintree{_lgraphqlSshendpointsbyprojectGraphql, map[string]*bintree{}},
		"updateDeployTarget.graphql":       &bintree{_lgraphqlUpdatedeploytargetGraphql, map[string]*bintree{}},
		"updateDeployTargetConfig.graphql": &bintree{_lgraphqlUpdatedeploytargetconfigGraphql, map[string]*bintree{}},
		"variables": &bintree{nil, map[string]*bintree{
			"addOrUpdateEnvVariableByName.graphql":            &bintree{_lgraphqlVariablesAddorupdateenvvariablebynameGraphql, map[string]*bintree{}},
			"deleteEnvVariableByName.graphql":                 &bintree{_lgraphqlVariablesDeleteenvvariablebynameGraphql, map[string]*bintree{}},
			"getEnvVariablesByProjectEnvironmentName.graphql": &bintree{_lgraphqlVariablesGetenvvariablesbyprojectenvironmentnameGraphql, map[string]*bintree{}},
		}},
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
