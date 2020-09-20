import SwiftUI
import NpnServer

@main
struct npnApp: App {
    init() {
        print("starting npn...")
        let path = NSSearchPathForDirectoriesInDomains(.libraryDirectory, .userDomainMask, true)
        let port = NpnServer.LibRun("ios", path[0])
        print("npn started on port [\(port)]")
        let url = URL.init(string: "http://localhost:\(port)/")!
        self.cv = ContentView(url: URLRequest(url: url))
    }

    var cv: ContentView
    
    var body: some Scene {
        WindowGroup {
            cv
        }
    }
}
