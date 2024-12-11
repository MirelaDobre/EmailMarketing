package main

func GetEmail(tip string) string {

	EmailCraciun := `
<!DOCTYPE html>
<html lang="ro">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Oferta de Crăciun</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      background-color: #f4f4f4;
      margin: 0;
      padding: 0;
    }
    .email-container {
      max-width: 600px;
      margin: 20px auto;
      background-color: #ffffff;
      border-radius: 8px;
      box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
      overflow: hidden;
    }
    .header {
      background-color: #d32f2f;
      color: white;
      text-align: center;
      padding: 20px 10px;
    }
    .header h1 {
      margin: 0;
      font-size: 24px;
    }
    .body {
      padding: 20px;
      text-align: center;
    }
    .body p {
      font-size: 16px;
      color: #333333;
      line-height: 1.5;
    }
    .cta {
      margin-top: 20px;
    }
    .cta a {
      background-color: #d32f2f;
      color: white;
      text-decoration: none;
      padding: 12px 20px;
      font-size: 16px;
      border-radius: 5px;
    }
    .cta a:hover {
      background-color: #b71c1c;
    }
    .footer {
      background-color: #f4f4f4;
      color: #666666;
      text-align: center;
      padding: 10px;
      font-size: 14px;
    }
    .footer a {
      color: #d32f2f;
      text-decoration: none;
    }
  </style>
</head>
<body>
  <div class="email-container">
    <div class="header">
      <h1>🎄 Oferta de Crăciun 🎁</h1>
    </div>
    <div class="body">
      <svg xmlns="http://www.w3.org/2000/svg" height="100px" viewBox="0 0 64 64" width="100px">
        <circle cx="32" cy="32" r="32" fill="#d32f2f" />
        <path fill="#fff" d="M32 14l8.5 17.2L58 36l-14.3 14 3.4 20.3L32 52 19 70l3.4-20L8 36l17.6-4.8L32 14z" />
      </svg>
      <p>
        Bucură-te de sărbători cu oferta noastră specială! Cumpără o cană de cafea de Crăciun și primești 20% reducere.
      </p>
      <div class="cta">
        <a href="https://www.exemplu.ro/oferta-craciun" target="_blank">Vezi Oferta</a>
      </div>
    </div>
    <div class="footer">
      <p>Mulțumim că sunteți alături de noi!</p>
      <p>Dacă dorești să te dezabonezi, fă click <a href="https://www.exemplu.ro/unsubscribe">aici</a>.</p>
    </div>
  </div>
</body>
</html>
`

	EmailOferte := `
<!DOCTYPE html>
<html lang="ro">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Oferta După Sărbători</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      background-color: #f4f4f4;
      margin: 0;
      padding: 0;
    }
    .email-container {
      max-width: 600px;
      margin: 20px auto;
      background-color: #ffffff;
      border-radius: 8px;
      box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
      overflow: hidden;
    }
    .header {
      background-color: #4caf50;
      color: white;
      text-align: center;
      padding: 20px 10px;
    }
    .header h1 {
      margin: 0;
      font-size: 24px;
    }
    .body {
      padding: 20px;
      text-align: center;
    }
    .body p {
      font-size: 16px;
      color: #333333;
      line-height: 1.5;
    }
    .cta {
      margin-top: 20px;
    }
    .cta a {
      background-color: #4caf50;
      color: white;
      text-decoration: none;
      padding: 12px 20px;
      font-size: 16px;
      border-radius: 5px;
    }
    .cta a:hover {
      background-color: #388e3c;
    }
    .footer {
      background-color: #f4f4f4;
      color: #666666;
      text-align: center;
      padding: 10px;
      font-size: 14px;
    }
    .footer a {
      color: #4caf50;
      text-decoration: none;
    }
  </style>
</head>
<body>
  <div class="email-container">
    <div class="header">
      <h1>🎉 Oferta Specială Generală 🎉</h1>
    </div>
    <div class="body">
      <svg xmlns="http://www.w3.org/2000/svg" height="100px" viewBox="0 0 64 64" width="100px">
        <circle cx="32" cy="32" r="32" fill="#4caf50" />
        <path fill="#fff" d="M32 10l10 20h20l-16 16 6 20L32 50l-20 16 6-20L2 30h20L32 10z" />
      </svg>
      <p>
        Începe noul an cu o ofertă deosebită! Descoperă reduceri de până la 30% la produsele noastre preferate.
      </p>
      <div class="cta">
        <a href="https://www.exemplu.ro/oferta-dupa-sarbatori" target="_blank">Descoperă Reducerile</a>
      </div>
    </div>
    <div class="footer">
      <p>Rămâi conectat pentru mai multe oferte!</p>
      <p>Dacă dorești să te dezabonezi, fă click <a href="https://www.exemplu.ro/unsubscribe">aici</a>.</p>
    </div>
  </div>
</body>
</html>`

	if tip == "craciun" {
		return EmailCraciun
	} else if tip == "oferte" {
		return EmailOferte
	} else {
		return ""
	}
}
