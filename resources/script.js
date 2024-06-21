fetch("http://localhost:8080/regenerate")

const URL = "http://localhost:8080/try"
const userInput = document.getElementById("userWordbox");
const userWords = document.getElementById("userWords");

win = false

function onBtnClick() {
    if (win) {
        alert("Вы выиграли! Обновите страницу!")
        return
    }

    var userInp = userInput.value.toUpperCase();

    fetch(URL, {
        method: 'POST',
        body: userInp
    })
    .then(r => {
        r.text().then(t => {
            if (!isJsonString(t)) {
                alert(t)
                return
            }
            js = JSON.parse(t)

            var buff = '';
            gg = true;
            for (let i = 0; i <= 5; i++) {
                let e = js[i]

                if (e != 0)
                    gg = false

                buff += '<span style=\"color:' + (e == 2 ? "red" : e == 1 ? "yellow" : "green") + '\">' + userInp[i] + '</span>'
            }
            userWords.innerHTML += "<li>" + buff + "</li>"

            if (gg == true) {
                win = true
                alert("Вы выиграли! Обновите страницу!")
            }
        })
    })
}

function isJsonString(str) {
    try {
        JSON.parse(str);
    } catch (e) {
        return false;
    }
    return true;
}