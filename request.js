const axios = require('axios');
const hotpOtpGenerator = require('hotp-totp-generator');

const mail = 'brandonma98@protonmail.com';
const body = {
  github_url:
    'https://gist.github.com/BrandonMA/abe33fc36d693c1edfd4eb6be38daeaf',
  contact_email: mail,
};

const totpToken = hotpOtpGenerator.totp({
  key: `${mail}HENNGECHALLENGE003`,
  algorithm: 'sha512',
  digits: 10,
});

const base64Data = Buffer.from(`${mail}:${totpToken}`).toString('base64');

const headers = {
  'Content-Type': 'application/json',
  Authorization: `Basic ${base64Data}`,
};

axios
  .post('https://api.challenge.hennge.com/challenges/003', body, {
    headers,
  })
  .then((response) => {
    console.log('You sent the application');
    console.log(response);
    console.log(response.data);
  })
  .catch((error) => {
    console.log(error);
  });
