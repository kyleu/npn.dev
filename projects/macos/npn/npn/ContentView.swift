import SwiftUI

struct ContentView: View {
    let u: URLRequest
    
    var body: some View {
        GeometryReader { g in
            ScrollView {
                WebView(url: self.u).frame(height: g.size.height)
            }.frame(height: g.size.height)
        }
    }

    init(url: URLRequest) {
        self.u = url
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView(url: URLRequest(url: URL.init(string: "https://localhost:10101")!))
    }
}
