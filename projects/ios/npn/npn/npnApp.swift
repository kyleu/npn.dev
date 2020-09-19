import SwiftUI
import NpnServer

@main
struct npnApp: App {
    var port: Int32 = 0
    
    init() {
        print("starting npn...")
        port = NpnServer.LibRun()
        print("npn started on port [\(port)]")
    }
    
    var body: some Scene {
        WindowGroup {
            ContentView()
        }
    }
}
