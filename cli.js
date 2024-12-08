#!/usr/bin/env node
const { execFileSync } = require('child_process');
const os = require('os');
const path = require('path');

// Detectar sistema operacional
const platform = os.platform();
let binary = null;

if (platform === 'win32') {
  binary = path.join(__dirname, 'bin', 'gapizer.exe');
} else {
  console.error(`Sistema operacional não suportado: ${platform}`);
  process.exit(1);
}

try {
  // Passar argumentos do usuário para o binário
  execFileSync(binary, process.argv.slice(2), { stdio: 'inherit' });
} catch (err) {
  console.error(`Erro ao executar GAPIzer: ${err.message}`);
  process.exit(1);
}
