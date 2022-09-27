[Access content in English](README.md) 

<p align="center">
  <img src="images/secDevLabs-logo.png" allign="center" height=""/>
  <!-- logo font: Agency FB Bold Condensed -->
</p>

<p align="center">
Laboratório para aprender desenvolvimento web e mobile seguro de forma prática.
</p>

<p align="center">
<a href="https://github.com/globocom/secDevLabs/blob/master/docs/CONTRIBUTING.md"><img src="https://img.shields.io/badge/PRs-Welcome-brightgreen"/></a>
<a href="https://gitter.im/secDevLabs/community"><img src="https://badges.gitter.im/secDevLabs/community.svg"/></a>
</p>

## Construa seu ambiente 

Ao usar esse laborátorio em um ambiente via docker-compose, você aprenderá como os riscos de segurança de aplicativos da Web mais críticos são explorados e como esses códigos vulneráveis podem ser corrigidos para não ter ameaças ao seu aplicativo. 👩‍💻

## Por onde começar? 

Depois de dar "forking" neste repositório, você encontrará vários aplicativos vulneráveis pretendidos com base em cenários da vida real em várias linguagens, como Golang, Python e PHP. Um bom começo seria instalar aqueles com os quais você está mais familiarizado. Você pode encontrar instruções para fazer isso em cada um dos aplicativos. 💡

Cada um desses aplicativos tem uma seção `Narrativa de ataque` que descreve como um invasor exploraria a vulnerabilidade correspondente. Antes de procurar a vulnerabilidade ou ler qualquer código, pode ser uma boa ideia seguir as etapas da seção 'Narativa de ataque' para que você possa entender melhor o ataque. 💉

Agora é hora de proteger o aplicativo! Imagine que este é o seu aplicativo e você precisa corrigir essas falhas! Sua missão é escrever novos códigos que eliminem as falhas encontradas e enviar um novo Pull Request para implantar um aplicativo seguro! 🔐

## Quão seguro é o meu novo código?

Depois de corrigir uma vulnerabilidade, você pode enviar um Pull Request para solicitar gentilmente à comunidade secDevLabs que revise seus novos códigos seguros. Se você está se sentindo um pouco perdido, tente dar uma olhada nas soluções já enviadas, pode ser que elas o ajudem! 🚀

## OWASP Top 10 (2021) apps: 💻

Isenção de responsabilidade: você está prestes a instalar aplicativos vulneráveis em sua máquina! 🔥

| Vulnerability                                       | Language       | Application                                                                    |
| ----------------------------------------------------| -------------- | ------------------------------------------------------------------------------ |
| A1 - Quebra de controle de acesso                   | Golang         | [Vulnerable Ecommerce API](owasp-top10-2021-apps/a1/ecommerce-api)             |
| A1 - Quebra de controle de acesso                   | NodeJS         | [Tic-Tac-Toe](owasp-top10-2021-apps/a1/tictactoe)                              |
| A1 - Quebra de controle de acesso                   | Golang         | [Camplake-API](owasp-top10-2021-apps/a1/camplake-api)                          |
| A2 - Falhas criptográficas.                         | Golang         | [SnakePro](owasp-top10-2021-apps/a2/snake-pro)                                 |
| A3 - Injeção                                        | Golang         | [CopyNPaste API](owasp-top10-2021-apps/a3/copy-n-paste)                        |
| A3 - Injeção                                        | NodeJS         | [Mongection](owasp-top10-2021-apps/a3/mongection)                              |
| A3 - Injeção                                        | Python         | [SSType](owasp-top10-2021-apps/a3/sstype)                                      |
| A3 - Injeção (XSS)                                  | Python         | [Gossip World](owasp-top10-2021-apps/a3/gossip-world)                          |
| A3 - Injeção (XSS)                                  | React          | [Comment Killer](owasp-top10-2021-apps/a3/comment-killer)                      |
| A3 - Injeção (XSS)                                  | Angular/Spring | [Streaming](owasp-top10-2021-apps/a3/streaming)                                |
| A5 - Configuração Insegura (XXE)                    | PHP            | [ViniJr Blog](owasp-top10-2021-apps/a5/vinijr-blog)                            |
| A5 - Configuração Insegura                          | PHP            | [Vulnerable Wordpress Misconfig](owasp-top10-2021-apps/a5/misconfig-wordpress) |
| A5 - Configuração Insegura                          | NodeJS         | [Stegonography](owasp-top10-2021-apps/a5/stegonography)                        |
| A6 - Componentes desatualizados e vulneráveis       | PHP            | [Cimentech](owasp-top10-2021-apps/a6/cimentech)                                |
| A6 - Componentes desatualizados e vulneráveis       | Python         | [Golden Hat Society](owasp-top10-2021-apps/a6/golden-hat)                      |
| A7 - Falha de Identificação e Autenticação          | Python         | [Saidajaula Monster Fit](owasp-top10-2021-apps/a7/saidajaula-monster)          |
| A7 - Falha de Identificação e Autenticação          | Golang         | [Insecure go project](owasp-top10-2021-apps/a7/insecure-go-project)            |
| A8 - Falha na Integridade de Dados e Software       | Python         | [Amarelo Designs](owasp-top10-2021-apps/a8/amarelo-designs)                    |
| A9 - Monitoramento de Falhas e Registros de Segurança| Python         | [GamesIrados.com](owasp-top10-2021-apps/a9/games-irados)                       |

## OWASP Top 10 (2016) Mobile apps: 📲

Isenção de responsabilidade: você está prestes a instalar aplicativos vulneráveis em sua máquina! 🔥

| Vulnerability                            | Language     | Application                                         |
| ---------------------------------------- | ------------ | --------------------------------------------------- |
| M2 - Armazenamento de dados inseguro     | Dart/Flutter | [Cool Games](owasp-top10-2016-mobile/m2/cool_games) |
| M4 - Autenticação insegura               | Dart/Flutter | [Note Box](owasp-top10-2016-mobile/m4/note-box)     |
| M5 - Criptografia insuficiente           | Dart/Flutter | [Panda Zap](owasp-top10-2016-mobile/m5/panda_zap)   |

## Contribuindo

Nós encorajamos você a contribuir com o SecDevLabs! Consulte a seção [Contribuindo para o SecDevLabs](/docs/CONTRIBUTING.md) para obter orientações sobre como contribuir! 🎉

## Licença

Este projeto está licenciado sob a Licença BSD 3-Clause 'Nova' ou 'Revisada' - leia o arquivo LICENSE.md para obter detalhes.📖
