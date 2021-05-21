import firebase from "firebase";
import React from "react";
import ReactDOM from "react-dom";

const PhraGen: React.FC = () => {
  const [phrase, setPhrase] = React.useState<string>("");

  React.useEffect(() => {
    const newPhrase = generate();
    setPhrase(newPhrase);
  });

  const onClick = (e: React.SyntheticEvent<HTMLButtonElement>) => {
    const target = e.currentTarget;
    switch (target.name) {
      case "regenerate":
        const newPhrase = generate();
        setPhrase(newPhrase);
        break;
      case "copyRaw":
        copyRaw(phrase);
        break;
      case "copyWithoutWhiteSpace":
        copyWithoutWhiteSpace(phrase);
        break;
      default:
        break;
    }
  };

  return (
    <div>
      <div className="content">
        <button
          className="button is-primary is-medium"
          name="regenerate"
          onClick={onClick}
        >
          Push to re-generate a phrase
        </button>
      </div>
      <div className="content">
        <div style={{ fontSize: "large" }}>generated phrase:</div>
        <div style={{ fontSize: "x-large" }}>{phrase}</div>
      </div>
      <button
        className="button is-info is-medium"
        name="copyRaw"
        onClick={onClick}
      >
        Copy to clipboard
      </button>
      &nbsp;
      <button
        className="button is-info is-medium"
        name="copyWithoutWhiteSpace"
        onClick={onClick}
      >
        Copy to clipboard (without whitespace)
      </button>
    </div>
  );
};

const generate = (): string => {
  firebase
    .app()
    .functions("asia-northeast1")
    .httpsCallable("generate")({})
    .then((response) => {
      if (response.data) {
        return response.data() as string;
      } else {
        throw new Error();
      }
    })
    .then((text) => {
      return text;
    })
    .catch((_) => {
      throw new Error();
    });
  return "";
};

const copyRaw = (str: string) => {
  copyText(str);
};

const copyWithoutWhiteSpace = (str: string) => {
  copyText(str.replace(/\s+/g, ""));
};

const copyText = (str: string) => {
  const tmp = document.createElement("div");
  tmp.appendChild(document.createElement("pre")).textContent = str;

  const s = tmp.style;
  s.position = "fixed";
  s.left = "-100%";

  document.body.appendChild(tmp);
  document.getSelection()?.selectAllChildren(tmp);
  document.execCommand("copy");
  document.body.removeChild(tmp);
};

ReactDOM.render(<PhraGen />, document.getElementById("phrase"));
