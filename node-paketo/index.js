const express=require('express')
const app=express()
const port=8080
app.get('/',(req,res)=>{
    res.send('I am just NodeJS Application what about you?')
})

app.listen(port,()=>{
    console.log(`Application is running on ${port}`)
})
