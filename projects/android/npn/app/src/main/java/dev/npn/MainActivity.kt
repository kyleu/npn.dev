package dev.npn

import android.os.Bundle
import android.util.Log
import android.webkit.WebView
import android.webkit.WebViewClient
import androidx.appcompat.app.AppCompatActivity

class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        Log.i("npn", "npn is starting")
        val path = getFilesDir().getAbsolutePath()
        val port = library.Library.run("android", path)
        Log.i("npn", "npn has started with path [" + path + "] on port [${port}]")
        setContentView(R.layout.activity_main)

        val webView: WebView = findViewById(R.id.webview)
        webView.setWebViewClient(WebViewClient())
        val settings = webView.getSettings();

        settings.loadsImagesAutomatically = true;
        settings.javaScriptEnabled = true;
        settings.domStorageEnabled = true;

        webView.loadUrl("http://localhost:${port}/")
    }
}
