// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
// import { getAnalytics } from "firebase/analytics";
// import { getAuth } from "firebase/auth";
import { getAuth, GoogleAuthProvider } from 'firebase/auth';
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyCk6foFZ6dS_jcpdtGy3yh8aq9GaTGgfpc",
  authDomain: "pengin-party.firebaseapp.com",
  projectId: "pengin-party",
  storageBucket: "pengin-party.firebasestorage.app",
  messagingSenderId: "826357668006",
  appId: "1:826357668006:web:c57cd447eb01c9d303fc67",
  measurementId: "G-ZZ9WHL4Z0D"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
// const analytics = getAnalytics(app);

// Initialize Firebase Authentication and get a reference to the service
const auth = getAuth(app);


// Google認証プロバイダの準備
const googleProvider = new GoogleAuthProvider();

export { auth, googleProvider };