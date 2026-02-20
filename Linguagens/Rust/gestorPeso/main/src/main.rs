fn main() {
    let nome = String::from("Richardt");
    let idade: u64 = 20;
    let mut peso: f32 = 82.5;
    let treinando: bool = true;
    //para declaração de variaveis e basicamente assim
    // let x: tipo : valor ou //let x= tipo ::from valor
    //declaar string
    println!(
        "Meu nome e {} atualmente {} malhando, meu peso atual{}",
        nome, idade, peso
    );
    peso = peso - 1.5;
    println!(
        "Após uma semana consegui perder {:?} o esquivalente a 1.5kg",
        peso
    );
}
