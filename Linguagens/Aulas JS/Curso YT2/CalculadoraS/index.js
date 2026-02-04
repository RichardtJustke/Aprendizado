let num1 = 10 ;
let num2 = 3;


console.log("Escolha uma opcao:");
console.log("1 - Adicao ➕");
console.log("2 - Subtracao ➖");
console.log("3 - Divisao ➗");
console.log("4 - Multiplicacao ✖️");
console.log("5 - Potencia ⚡");
console.log("6 - Fatorial da soma dos numeros 🧮");
console.log("7 - Tabuada da soma dos numeros (1 a 10) 📚");

let calc = 1;

switch (calc) {
	case 1:{
        console.log(num1 + num2)
        break;
    }
	case 2:{
        console.log(num1 - num2)
        break;
    }
	case 3:{
        console.log(num1 / num2)
        break;
    }
	case 4:{
        console.log(num1 * num2)
        break;
    }
	case 5:{
        console.log(num1 ** num2)
        break;
    }
	case 6:{
        let soma = num1 + num2
        let fatorial = 1 
        for(let i = 1; i <= soma; i++){
            fatorial *= i
        }   
        console.log(fatorial)
        break;
    }
	case 7:{
        let soma = num1 + num2
        for(let i = 1; i <= 10; i++){
            console.log(`${soma} x ${i} = ${soma * i}`)
        }
        break;
    }
	default:{

    }
}



