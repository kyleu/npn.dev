package cloudbeaver

var gqlOpenSession = `mutation OpenSession {
	session: openSession {
		createTime
		lastAccessTime
		cacheExpired
		locale
		connections {
			id
			name
			driverId
			connected
			features
		}
	}
}`

func gqlNavChildren(path string) string {
	return `query NavChildren {
		navNodeChildren(parentPath: "` + path + `") {
			id
			name
			description
			hasChildren
			nodeType
			icon
			folder
			inline
			navigable
			features
		}
	}`
}
