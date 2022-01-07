import {ColorModeScript} from "@chakra-ui/react"
import * as React from "react"
import ReactDOM from "react-dom"
import {App} from "./App"
import {initializeFirebase} from "./analytics/firebase";

// Redirect to custom URL if coming from Firebase Hosting provided URLs
if (window.location.hostname.indexOf((new URL(process.env.REACT_APP_DOMAIN!)).hostname) === -1) {
  window.location.replace(process.env.REACT_APP_DOMAIN!);
}

// Disable logs in production
if (process.env.NODE_ENV === 'production') {
  console.log = () => {}
  console.error = () => {}
  console.debug = () => {}
}

initializeFirebase()

ReactDOM.render(
  <React.StrictMode>
    <ColorModeScript/>
    <App/>
  </React.StrictMode>,
  document.getElementById("root"),
)
