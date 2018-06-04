import React    from 'react';
import ReactDOM from 'react-dom';

class PhraGen extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      phrase: ""
    }
  }

  componentDidMount() {
    this.generate();
  }

  generate() {
    fetch('/api/v1/phrase', {
      method: 'GET'
    }).then(response => {
      if (response.ok) {
        return response.text();
      } else {
        throw new Error();
      }
    }).then(text => {
      this.setState({phrase:text});
    }).catch(error => {
      console.log(error);
    });
  }

  copyText(str) {
    const tmp = document.createElement('div');
    tmp.appendChild(document.createElement('pre')).textContent = str;

    const s = tmp.style;
    s.position = 'fixed';
    s.left = '-100%';

    document.body.appendChild(tmp);
    document.getSelection().selectAllChildren(tmp);
    document.execCommand('copy');
    document.body.removeChild(tmp);
  }

  copyRaw(id) {
    const text = this.state.phrase
    this.copyText(text);
  }

  copyWithoutWhiteSpace(id) {
    const text = this.state.phrase
    this.copyText(text.replace(/\s+/g, ""));
  }

  render() {
    return (
    <div>
      <div>
        <span><button onClick={this.generate.bind(this)}>Push to re-generate a phrase</button></span>
        <span><button onClick={this.copyRaw.bind(this)}>Copy to clipboard</button></span>
        <span><button onClick={this.copyWithoutWhiteSpace.bind(this)}>Copy to clipboard (without whitespace)</button></span>
      </div>
      <div style={{fontSize:"x-large"}}>
        generated phrase:
      </div>
      <div style={{fontSize:"x-large"}}>
        {this.state.phrase}
      </div>
    </div>
    )
  }
}

ReactDOM.render(
  <PhraGen />,
  document.getElementById('phrase')
);

