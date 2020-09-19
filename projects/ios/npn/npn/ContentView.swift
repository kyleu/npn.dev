import SwiftUI

struct ContentView: View {
    @ObservedObject var webViewStateModel: WebViewStateModel = WebViewStateModel()
    
    var body: some View {
        WebView(url: URL.init(string: "http://localhost:10101")!, webViewStateModel: self.webViewStateModel, onNavigationAction: nil)
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
