import "@/infrastructure/firebase.ts"
import { createUserWithEmailAndPassword } from "firebase/auth";
import { auth } from "@/infrastructure/firebase";

const FirebaseAuthTest = () => {
    const email = "test@example.com";
    const password = "password";

    // const auth = getAuth();
    createUserWithEmailAndPassword(auth, email, password)
        .then((userCredential) => {
            // Signed up 
            const user = userCredential.user;
            const providerId = userCredential.providerId;
            const operationType = userCredential.operationType;

            console.log("user: ");
            console.log(user);
            console.log("providerId: " + providerId);
            console.log("operationType: ");
            console.log(operationType);
        })
        .catch((error) => {
            const errorCode = error.code;
            const errorMessage = error.message;

            console.error("errorCode: " + errorCode);
            console.error("errorMessage: " + errorMessage);
        });
    return (
        <div className="test">あいうえお</div>
    );
}

export default FirebaseAuthTest