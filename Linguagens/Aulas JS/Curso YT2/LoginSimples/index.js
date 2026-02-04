let usuarioCorreto = "admin"
let senhaCorreta = "134"

function logIn(usuarioCorreto, senhaCorreta){
    if(usuarioCorreto === "admin" && senhaCorreta === "1234"){
        console.log('Login aprovado')
    }else{
        console.log('Usuário ou senha inválidos')
    }
}
console.log(logIn(usuarioCorreto, senhaCorreta))