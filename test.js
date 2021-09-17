
printHello().then(data=>{
    console.log("data is: ", data)
}).catch(e=>{
    console.log("Error", e)
})


let user ={
    name: "manish",
    // age: "30"
}

let {name, age}= user;

<button id="test" onClick></button>



function printHello(){
    return new Promise((resolve, reject)=>{

        setTimeout(()=>resolve("Hellow world ")
        ,1000)
    }
}



