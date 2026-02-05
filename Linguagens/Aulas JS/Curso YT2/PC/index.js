const computador = {
    processador: 'Ryzen',
    memoria: '16GB',
    armazenamento: '256GB',
    ligado: false,

    ligar() {
        this.ligado = true;
    },

    desligar() {
        this.ligado = false;
    }
};

// Mostrar estado antes
console.log('Antes de ligar:', computador);

// Ligar o computador
computador.ligar();
console.log('Depois de ligar:', computador);

// Desligar o computador
computador.desligar();
console.log('Depois de desligar:', computador);