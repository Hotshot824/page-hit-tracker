(() => {
    let url = new URL('{{ .PUBLISHING_URL }}');
    let params = { href: window.location.origin + window.location.pathname };
    Object.keys(params).forEach(key => url.searchParams.append(key, params[key]));
    fetch(url, {
        method: 'GET',
        cache: 'no-cache',
        credentials: 'omit',
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
    }).then(response => response.json())
        .then(data => {
            let count = data.count;
            document.getElementById("page-hit-tracker").innerHTML = count;
        })
        .catch(error => console.log(error));
})();