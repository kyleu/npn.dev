import SwiftUI
import WebKit

struct WebView: NSViewRepresentable {
    let view: WKWebView = WKWebView()

    let request: URLRequest

    func makeNSView(context: Context) -> WKWebView {
        view.load(request)
        view.allowsBackForwardNavigationGestures = true
        return view
    }

    func updateNSView(_ view: WKWebView, context: Context) {
        
    }
    
    init(url: URLRequest) {
        self.request = url
    }
}
