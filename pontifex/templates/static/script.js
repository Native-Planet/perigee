// disable perma-button class after its been clicked
document.addEventListener('DOMContentLoaded', function() {
    const actionButtons = document.querySelectorAll('.perma-button');
    actionButtons.forEach(button => {
        button.addEventListener('click', function(e) {
            if (!this.classList.contains('disabled')) {
                this.classList.add('disabled');
                const buttonId = this.getAttribute('data-button-id') || this.textContent.trim();
                localStorage.setItem(`button-${buttonId}`, 'disabled');
                if (!this.hasAttribute('hx-post')) {
                    e.preventDefault();
                }
            }
        });
    });
    actionButtons.forEach(button => {
        const buttonId = button.getAttribute('data-button-id') || button.textContent.trim();
        if (localStorage.getItem(`button-${buttonId}`) === 'disabled') {
            button.classList.add('disabled');
        }
    });
});


// walletconnect stuff
async function connectWallet() {
    const provider = await WalletConnectProvider.init({
        projectId: '0d7e141c40636bd870f77fcc6eb513b9',
        chains: [1] // mainnet
    });
    const accounts = await provider.enable();
    htmx.ajax('POST', '/connect-wallet', {
        values: { address: accounts[0] }
    });
}

async function signMessage(message) {
    const signature = await provider.request({
        method: 'personal_sign',
        params: [message, account]
    });
    htmx.ajax('POST', '/verify-signature', {
        values: { signature }
    });
}

async function initWC() {
    if (!window.ethereum) {
        document.querySelector('.wallet-status').innerHTML = 
            '<span>DEVICE STATUS: Not connected</span>';
        return;
    }
    try {
        await window.ethereum.request({ method: 'eth_requestAccounts' });
        const accounts = await window.ethereum.request({ 
            method: 'eth_accounts' 
        });
        if (accounts[0]) {
            document.querySelector('.wallet-status').innerHTML = 
                '<span>DEVICE STATUS: Connected [' + 
                accounts[0].slice(0,6) + '...' + 
                accounts[0].slice(-4) + ']</span>';
            document.querySelector('.wallet-status').classList.remove('disconnected');
            document.querySelector('.wallet-status').classList.add('connected');
            document.getElementById('wallet-address').value = accounts[0];
            document.getElementById('submit-hw').style.display = 'block';
        }
    } catch (err) {
        document.querySelector('.wallet-status').innerHTML = 
            '<span>DEVICE STATUS: ERROR - ' + err.message + '</span>';
    }
}

async function signChallenge() {
    const challenge = document.getElementById('challenge').value;
    const address = document.getElementById('address').value;
    const ship = document.getElementById('ship').value;
    try {
        const signature = await window.ethereum.request({
            method: 'personal_sign',
            params: [challenge, address]
        });
        htmx.ajax('POST', '/auth', {
            values: {
                ship: ship,
                address: address,
                signature: signature,
                challenge: challenge,
                auth_type: 'hardware'
            },
            target: '#main-content',
            targetError: '#login-error'
        });
    } catch (err) {
        htmx.ajax('POST', '/auth/error', {
            values: { error: err.message },
            target: '#login-error'
        });
    }
}