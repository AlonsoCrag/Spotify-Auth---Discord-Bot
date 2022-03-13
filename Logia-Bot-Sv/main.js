const express = require('express');
const app = express();
const axios = require('axios').default

const send_code = async code => {
    await axios.get(`http://localhost:4000/accesstoken?code=${code}`);
}

const Application = {
    init: function () {
        app.get('/access/', (req, resp) => {
            console.log(req.query)
            resp.send(req.query.code)
            send_code(req.query.code)
            .then(() => console.log("Access token was sent to the Discord Server Bot"))
            .catch(err => console.log("Unexpected error"))
        });

        app.listen(3000, () => console.log("Listening at port 3000"));
    }
}

Application.init()