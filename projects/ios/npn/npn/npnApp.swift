import SwiftUI
import NpnServer

@main
struct npnApp: App {
    var port: Int32 = 0

    init() {
        print("starting npn...")
        let path = NSSearchPathForDirectoriesInDomains(.libraryDirectory, .userDomainMask, true)
        port = NpnServer.LibRun(path[0])
        print("npn started on port [\(port)]")
    }

    var body: some Scene {
        WindowGroup {
            ContentView()
        }
    }
}
