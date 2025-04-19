const express = require('express');


const app = express();
const port = 8080;

const cors = require('cors');

app.use(cors());
app.use(express.json());
app.use(express.urlencoded({extended:true}))


app.get("/", (req,res) =>{
    res.status(200).json({message:"Hello from server"})
});


app.post("/api", (req,res) => {
    let myJson = req.body;

    res.status(200).send(myJson);
});


app.post("/platform", (req,res) => {
    res.status(200).send(JSON.stringify(req.body));
})



app.listen(port, ()=>{
    console.log(`Server is running on port ${port}`)
})