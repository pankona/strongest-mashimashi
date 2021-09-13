import * as firebaseApp from "firebase/app";
import * as firebaseFunctions from "firebase/functions";

export const app = {
  initializeApp: () =>
    firebaseApp.initializeApp({
      apiKey: "AIzaSyCyGzqeB5UYGmU_qsiMRAcZa88kwEezCD0",
      authDomain: "strongest-mashimashi.firebaseapp.com",
      projectId: "strongest-mashimashi",
      storageBucket: "strongest-mashimashi.appspot.com",
      messagingSenderId: "404083950358",
      appId: "1:404083950358:web:861de5f1597aef9e604637",
      measurementId: "G-76Y4K0PXX8",
    }),
};
export const functions = {
  generate: async (): Promise<string> => {
    const response: any = await firebaseFunctions.httpsCallable(
      firebaseFunctions.getFunctions(),
      "generate"
    )({});
    return await response.data.phrase;
  },
};
