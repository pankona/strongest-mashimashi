import * as React from "react";
import * as ReactDOM from "react-dom";

interface IState {
  phrase: string;
}

class PhraGen extends React.Component<{}, IState> {
  constructor(props) {
    super(props);
    this.state = {
      phrase: ""
    };
    this.onClick = this.onClick.bind(this);
  }

  public componentDidMount() {
    this.generate();
  }

  public render() {
    return (
      <div>
        <div className="content">
          <button className="button is-primary is-medium" name="regenerate" onClick={this.onClick}>
            Push to re-generate a phrase
          </button>
        </div>
        <div className="content">
          <div style={{ fontSize: "large" }}>generated phrase:</div>
          <div style={{ fontSize: "x-large" }}>{this.state.phrase}</div>
        </div>
        <button className="button is-info is-medium" name="copyRaw" onClick={this.onClick}>
          Copy to clipboard
        </button>&nbsp;
        <button className="button is-info is-medium" name="copyWithoutWhiteSpace" onClick={this.onClick}>
          Copy to clipboard (without whitespace)
        </button>
      </div>
    );
  }

  private onClick(e: React.SyntheticEvent<HTMLButtonElement>) {
    const target = e.currentTarget;
    switch (target.name) {
      case "regenerate":
        this.generate();
        break;
      case "copyRaw":
        this.copyRaw();
        break;
      case "copyWithoutWhiteSpace":
        this.copyWithoutWhiteSpace();
        break;
      default:
        break;
    }
  }

  private copyRaw() {
    const text = this.state.phrase;
    this.copyText(text);
  }

  private copyWithoutWhiteSpace() {
    const text = this.state.phrase;
    this.copyText(text.replace(/\s+/g, ""));
  }

  private generate() {
    // this.setState({ phrase: "super simple phrase" });
    // return
    fetch("/api/v1/phrase", {
      method: "GET"
    })
      .then(response => {
        if (response.ok) {
          return response.text();
        } else {
          throw new Error();
        }
      })
      .then(text => {
        this.setState({ phrase: text });
      })
      .catch(error => {
        throw new Error();
      });
  }

  private copyText(str) {
    const tmp = document.createElement("div");
    tmp.appendChild(document.createElement("pre")).textContent = str;

    const s = tmp.style;
    s.position = "fixed";
    s.left = "-100%";

    document.body.appendChild(tmp);
    document.getSelection().selectAllChildren(tmp);
    document.execCommand("copy");
    document.body.removeChild(tmp);
  }
}

ReactDOM.render(<PhraGen />, document.getElementById("phrase"));
