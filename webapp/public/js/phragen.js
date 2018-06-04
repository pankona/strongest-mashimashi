
window.addEventListener('load', _ => {
    generate();
    document.getElementById('submit').addEventListener('click', _ => {
        generate();
    });
});

const generate = () => {
    fetch('/api/v1/phrase', {
        method: 'GET'
    }).then(response => {
        if (response.ok) {
            return response.text();
        } else {
            throw new Error();
        }
    }).then(text => {
        document.getElementById('phrase').textContent = text;
    }).catch(error => {
        console.log(error);
    });
}

const copyText = (str) => {
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

const copyRaw = (id) => {
    const text = document.getElementById(id).textContent
    if (text === '--- --- ---') {
        return
    }
    copyText(text);
}

const copyWithoutWhiteSpace = (id) => {
    const text = document.getElementById(id).textContent
    if (text === '--- --- ---') {
        return
    }
    copyText(text.replace(/\s+/g, ""));
}

