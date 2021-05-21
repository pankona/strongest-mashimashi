import React from "react";
import ReactDOM from "react-dom";
import firebase from "./firebase";

const PhraGen: React.FC = (): JSX.Element => {
  const [phrase, setPhrase] = React.useState<string>("");

  const fetchPhrase = () => {
    generate()
      .then((newPhrase: string) => {
        setPhrase(newPhrase);
      })
      .catch((err) => {
        throw new Error(err);
      });
  };

  React.useEffect(() => {
    fetchPhrase();
  }, []);

  const onClick = (e: React.SyntheticEvent<HTMLButtonElement>) => {
    const target = e.currentTarget;
    switch (target.name) {
      case "regenerate":
        fetchPhrase();
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

const generate = async (): Promise<string> => {
  const phrase = await firebase
    .app()
    .functions("asia-northeast1")
    .httpsCallable("generate")({})
    .then((response) => {
      if (response.data) {
        return response.data.phrase;
      } else {
        throw new Error();
      }
    })
    .then((text) => {
      return text;
    })
    .catch((err) => {
      throw new Error(err);
    });
  return phrase;
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
