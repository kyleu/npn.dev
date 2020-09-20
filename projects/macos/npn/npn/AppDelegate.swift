import Cocoa
import SwiftUI

@NSApplicationMain
class AppDelegate: NSObject, NSApplicationDelegate {
    var window: NSWindow!
    var task: Process = Process()

    func applicationDidFinishLaunching(_ aNotification: Notification) {
        initProcess()
    }

    func applicationShouldTerminateAfterLastWindowClosed(_ sender: NSApplication) -> Bool {
        true
    }

    func applicationWillTerminate(_ notification: Notification) {
        task.terminate()
    }

    func initProcess() {
        let path = Bundle.main.path(forResource: "npn-server", ofType: "")
        task.launchPath = path
        task.arguments = ["-p", "0"]

        let pipe = Pipe()
        task.standardOutput = pipe
        let outHandle = pipe.fileHandleForReading
        outHandle.waitForDataInBackgroundAndNotify()

        var progressObserver : NSObjectProtocol!
        progressObserver = NotificationCenter.default.addObserver(forName: NSNotification.Name.NSFileHandleDataAvailable, object: outHandle, queue: nil) {
            notification -> Void in
            let data = outHandle.availableData

            if let str = String(data: data, encoding: String.Encoding.utf8) {
                if let range: Range<String.Index> = str.range(of: "port:") {
                    let newline = str.firstIndex(of: "\n")!
                    let portString = String(str[range.upperBound..<newline]).trimmingCharacters(in: .whitespacesAndNewlines)
                    let port = Int(portString)!
                    // print("PORT: ", port)
                    self.initWindow(port: port)
                } else {
                    print("substring not found in \(str)")
                }
            }
            NotificationCenter.default.removeObserver(progressObserver!)
            // outHandle.waitForDataInBackgroundAndNotify()
        }

        task.launch()

        outHandle.waitForDataInBackgroundAndNotify()
    }

    func initWindow(port: Int) {
        let urlString = "http://localhost:\(port)"
        print(urlString)
        let url = URL(string: urlString)
        let request = URLRequest(url: url!)

        let contentView = ContentView(url: request)

        window = NSWindow(
            contentRect: NSRect(x: 0, y: 0, width: 1280, height: 720),
            styleMask: [.titled, .closable, .miniaturizable, .resizable, .fullSizeContentView],
            backing: .buffered,
            defer: false
        )
        window.isReleasedWhenClosed = false
        window.center()
        window.title = "npn"
        window.setFrameAutosaveName("npn")
        window.contentView = NSHostingView(rootView: contentView)
        window.makeKeyAndOrderFront(nil)
    }
}
