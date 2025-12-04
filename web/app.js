const statusEl = document.getElementById("status");
const contentEl = document.getElementById("main-content");
const outputEl = document.getElementById("action-output");

const state = {
    session: null,
    config: {
        rollerUrl: "https://roller.urbit.org",
        ethProvider: "",
    },
};

const go = new Go();

const wasmStartup = startWasm()
    .catch((err) => {
        // Keep status empty unless there's a failure
        setStatus(`Failed to load bundle: ${err}`, "error");
    });

renderLogin();

function setStatus(message, variant = "info") {
    if (variant === "error") {
        statusEl.style.display = "block";
        statusEl.textContent = message;
        statusEl.style.color = "#800000";
    } else {
        statusEl.textContent = "";
        statusEl.style.display = "none";
    }
}

function startWasm() {
    if (window.__perigeeWasmPromise) {
        return window.__perigeeWasmPromise;
    }
    window.__perigeeWasmPromise = (async () => {
        const response = await fetch("./perigee.wasm");
        if (!response.ok) {
            throw new Error(`fetch failed (${response.status})`);
        }
        const bytes = await response.arrayBuffer();
        const { instance } = await WebAssembly.instantiate(bytes, go.importObject);
        go.run(instance);
        await waitForPerigee();
        if (window.perigee?.setRollerURL) {
            try {
                await window.perigee.setRollerURL(state.config.rollerUrl);
                await window.perigee.setEthProvider(state.config.ethProvider);
            } catch (err) {
                console.warn("endpoint init failed", err);
            }
        }
    })();
    return window.__perigeeWasmPromise;
}

async function waitForPerigee() {
    const maxWaitMs = 15000;
    const start = Date.now();
    while (!window.perigee) {
        if (Date.now() - start > maxWaitMs) {
            throw new Error("timed out waiting for wasm exports");
        }
        await new Promise((resolve) => setTimeout(resolve, 20));
    }
}

async function call(method, ...args) {
    await wasmStartup;
    const api = window.perigee;
    if (!api || typeof api[method] !== "function") {
        throw new Error(`wasm method '${method}' unavailable`);
    }
    const result = await api[method](...args);
    if (typeof result === "string") {
        try {
            return JSON.parse(result);
        } catch (_) {
            return result;
        }
    }
    return result;
}

function renderLogin() {
    contentEl.innerHTML = `
        <div class="form-group">
            <span class="centered">
                <h4>Azimuth public key cryptography operation web portal</h4>
                <div class="divider">* * *</div>
            </span>
            <form class="login-form" id="login-form">
                <div class="form-group">
                    <label for="ship">SHIP NAME</label>
                    <input type="text" id="ship" name="ship" placeholder="~sampel-palnet" required>
                </div>
                <div class="form-group">
                    <label for="ticket">MASTER TICKET / ETH PRIVATE KEY</label>
                    <input type="password" id="ticket" name="ticket" required>
                </div>
                <div class="form-group">
                    <label class="checkbox-container">
                        <input type="checkbox" id="use-passphrase">
                        Use passphrase
                    </label>
                    <div id="passphrase-section" class="form-group" style="display: none">
                        <input type="password" id="passphrase" name="passphrase" placeholder="Optional passphrase">
                    </div>
                </div>
                <div class="centered">
                    <button class="action-button" type="submit">Submit</button>
                </div>
            </form>
            <div class="divider">* * *</div>
        </div>
        ${endpointForm()}
    `;
    wireEndpointForm();
    const usePassphrase = contentEl.querySelector("#use-passphrase");
    const passphraseSection = contentEl.querySelector("#passphrase-section");
    usePassphrase.addEventListener("change", () => {
        passphraseSection.style.display = usePassphrase.checked ? "block" : "none";
    });
    const form = contentEl.querySelector("#login-form");
    form.addEventListener("submit", async (event) => {
        event.preventDefault();
        setStatus("Authenticating...");
        const ship = form.ship.value.trim();
        const ticket = form.ticket.value.trim();
        const passphrase = usePassphrase.checked ? form.passphrase.value : "";
        try {
            const session = await call("login", ship, ticket, passphrase);
            state.session = session;
            setStatus(`Authenticated as ${session.ship}`, "ok");
            renderDashboard();
        } catch (err) {
            console.error(err);
            setStatus(err?.message || String(err), "error");
        }
    });
}

function renderDashboard() {
    if (!state.session) {
        renderLogin();
        return;
    }
    const session = state.session;
    const point = session.point || {};
    const p = point.point || {};
    const ownership = p.ownership || {};
    const network = p.network || {};
    const currentLife = network.keys ? network.keys.life : "";

    contentEl.innerHTML = `
        ${endpointForm()}
        <div class="content-section">
            <div class="action-box">
                <h3>Identity</h3>
                <div class="ship-header">
                    <div>
                        <h2 class="ship-name">${point.patp || session.ship}</h2>
                        <span class="ship-dominion">
                            <h4>${p.dominion || "unknown"} ${point.clan || ""}</h4>
                        </span>
                    </div>
                </div>
                <div class="info-section">
                    ${infoRow("Owner", ownership.owner?.address)}
                    ${infoRow("Management", ownership.managementProxy?.address)}
                    ${infoRow("Transfer proxy", ownership.transferProxy?.address)}
                    ${infoRow("Spawn proxy", ownership.spawnProxy?.address)}
                </div>
                ${session.authType === "ticket" ? downloadButtons() : ""}
            </div>

            <div class="action-box">
                <h3>Keys</h3>
                <div class="info-section">
                    <span>Network Keys (Life: ${currentLife || "?"}, Rift: ${network.rift || "?"})</span>
                    <div class="key-row">
                        <span>Authentication</span>
                        <div class="address-display">${network.keys?.auth || "n/a"}</div>
                    </div>
                    <div class="key-row">
                        <span>Encryption</span>
                        <div class="address-display">${network.keys?.crypt || "n/a"}</div>
                    </div>
                    <form id="breach-form" class="breach-form">
                        <div class="info-row">
                            <span>Seed (hex, optional)</span>
                            <div class="address-input-container">
                                <input type="text" id="breach-seed" placeholder="Required when authenticating with eth privkey">
                            </div>
                            <button type="submit" class="action-button warning">Breach</button>
                        </div>
                    </form>
                </div>
            </div>

            <div class="action-box">
                <h3>Sponsorship</h3>
                <div class="info-section">
                    <form id="escape-form" class="transfer-form">
                        <div class="info-row">
                            <span>New sponsor</span>
                            <div class="address-input-container">
                                <input type="text" id="escape-sponsor" placeholder="~sponsor" required>
                            </div>
                            <button type="submit" class="action-button warning">Escape</button>
                        </div>
                    </form>
                    <div class="transfer-buttons">
                        <button id="cancel-escape" class="action-button">Cancel escape</button>
                        <form id="adopt-form" class="transfer-form" style="width:100%">
                            <div class="info-row">
                                <span>Adopt</span>
                                <div class="address-input-container">
                                    <input type="text" id="adopt-child" placeholder="~child" required>
                                </div>
                                <button type="submit" class="action-button warning">Adopt</button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>

            <div class="action-box">
                <h3>Transfers</h3>
                <div class="info-section">
                    ${transferForm("owner", "New owner (will reset keys)", "0x...")}
                    ${transferForm("management", "Management proxy", "0x...")}
                    ${transferForm("transfer-proxy", "Transfer proxy", "0x...")}
                    ${transferForm("spawn-proxy", "Spawn proxy", "0x...")}
                    ${transferForm("voting-proxy", "Voting proxy", "0x...")}
                </div>
            </div>
        </div>
    `;

    wireEndpointForm();
    wireDashboardHandlers();
}

function endpointForm() {
    return `
        <div class="action-box">
            <h3>Endpoints</h3>
            <form id="endpoint-form" class="transfer-form">
                <div class="info-row">
                    <span>Roller URL</span>
                    <div class="address-input-container">
                        <input type="text" id="roller-url" placeholder="https://roller.urbit.org">
                    </div>
                </div>
                <div class="info-row">
                    <span>ETH provider</span>
                    <div class="address-input-container">
                        <input type="text" id="eth-provider" placeholder="https://mainnet.infura.io/v3/<key>">
                    </div>
                </div>
                <div class="transfer-buttons">
                    <button type="submit" class="action-button">Apply</button>
                </div>
            </form>
        </div>
    `;
}

function wireEndpointForm() {
    const form = contentEl.querySelector("#endpoint-form");
    if (!form) return;
    const rollerInput = form.querySelector("#roller-url");
    const ethInput = form.querySelector("#eth-provider");
    if (rollerInput && !rollerInput.value) {
        rollerInput.value = state.config.rollerUrl || "https://roller.urbit.org";
    }
    if (ethInput && !ethInput.value) {
        ethInput.value = state.config.ethProvider || "";
    }
    form.addEventListener("submit", async (event) => {
        event.preventDefault();
        const rollerUrl = (rollerInput?.value || "").trim() || "https://roller.urbit.org";
        const ethProvider = (ethInput?.value || "").trim();
        setStatus("Applying endpoints...");
        try {
            await call("setRollerURL", rollerUrl);
            await call("setEthProvider", ethProvider);
            state.config.rollerUrl = rollerUrl;
            state.config.ethProvider = ethProvider;
            setStatus("Endpoints applied.", "ok");
        } catch (err) {
            console.error(err);
            setStatus(err?.message || String(err), "error");
        }
    });
}

function infoRow(label, value) {
    const display = value || "n/a";
    return `
        <div class="info-row">
            <span>${label}</span>
            <div class="address-display">${display}</div>
            <div></div>
        </div>
    `;
}

function downloadButtons() {
    return `
        <div class="download-section">
            <button class="action-button" id="download-wallet">Download wallet</button>
            <button class="action-button" id="download-keyfile">Download keyfile</button>
        </div>
    `;
}

function transferForm(kind, label, placeholder) {
    const id = `transfer-${kind}`;
    return `
        <form id="${id}" class="transfer-form">
            <div class="info-row">
                <span>${label}</span>
                <div class="address-input-container">
                    <input type="text" placeholder="${placeholder}" required>
                </div>
                <button type="submit" class="action-button warning">Submit</button>
            </div>
        </form>
    `;
}

function wireDashboardHandlers() {
    const walletBtn = contentEl.querySelector("#download-wallet");
    if (walletBtn) {
        walletBtn.addEventListener("click", async () => {
            setStatus("Generating wallet...");
            try {
                const wallet = await call("wallet");
                downloadText(`${state.session.ship || "ship"}-wallet.json`, JSON.stringify(wallet, null, 2), "application/json");
                setStatus("Wallet generated.", "ok");
            } catch (err) {
                console.error(err);
                setStatus(err?.message || String(err), "error");
            }
        });
    }

    const keyfileBtn = contentEl.querySelector("#download-keyfile");
        if (keyfileBtn) {
            keyfileBtn.addEventListener("click", async () => {
                setStatus("Generating keyfile...");
            try {
                const keyfile = await call("keyfile");
                downloadText(`${state.session.ship || "ship"}.key`, keyfile, "text/plain");
                setStatus("Keyfile generated.", "ok");
            } catch (err) {
                console.error(err);
                setStatus(err?.message || String(err), "error");
            }
        });
    }

    const breachForm = contentEl.querySelector("#breach-form");
    breachForm?.addEventListener("submit", async (event) => {
        event.preventDefault();
        const seed = contentEl.querySelector("#breach-seed").value.trim();
        setStatus("Submitting breach...");
        try {
            const receipt = await call("breach", seed);
            showOutput("Breach receipt", receipt);
            setStatus("Breach submitted.", "ok");
        } catch (err) {
            console.error(err);
            setStatus(err?.message || String(err), "error");
        }
    });

    const escapeForm = contentEl.querySelector("#escape-form");
    escapeForm?.addEventListener("submit", async (event) => {
        event.preventDefault();
        const sponsor = contentEl.querySelector("#escape-sponsor").value.trim();
        setStatus("Submitting escape...");
        try {
            const receipt = await call("escape", sponsor);
            showOutput("Escape receipt", receipt);
            setStatus("Escape submitted.", "ok");
        } catch (err) {
            console.error(err);
            setStatus(err?.message || String(err), "error");
        }
    });

    const cancelEscape = contentEl.querySelector("#cancel-escape");
    cancelEscape?.addEventListener("click", async () => {
        setStatus("Cancelling escape...");
        try {
            const receipt = await call("cancelEscape");
            showOutput("Escape cancellation receipt", receipt);
            setStatus("Escape cancelled.", "ok");
        } catch (err) {
            console.error(err);
            setStatus(err?.message || String(err), "error");
        }
    });

    const adoptForm = contentEl.querySelector("#adopt-form");
    adoptForm?.addEventListener("submit", async (event) => {
        event.preventDefault();
        const adoptee = contentEl.querySelector("#adopt-child").value.trim();
        setStatus("Submitting adoption...");
        try {
            const receipt = await call("adopt", adoptee);
            showOutput("Adoption receipt", receipt);
            setStatus("Adoption submitted.", "ok");
        } catch (err) {
            console.error(err);
            setStatus(err?.message || String(err), "error");
        }
    });

    wireTransfer("owner", true);
    wireTransfer("management");
    wireTransfer("transfer-proxy");
    wireTransfer("spawn-proxy");
    wireTransfer("voting-proxy");
}

function wireTransfer(kind, includeReset = false) {
    const form = contentEl.querySelector(`#transfer-${kind}`);
    if (!form) return;
    form.addEventListener("submit", async (event) => {
        event.preventDefault();
        const input = form.querySelector("input");
        const address = input.value.trim();
        const reset = includeReset ? true : false;
        setStatus(`Submitting ${kind} transfer...`);
        try {
            const receipt = await call("transfer", kind, address, reset);
            showOutput(`${kind} transfer receipt`, receipt);
            setStatus("Transfer submitted.", "ok");
        } catch (err) {
            console.error(err);
            setStatus(err?.message || String(err), "error");
        }
    });
}

function showOutput(title, payload) {
    const printable = typeof payload === "string" ? payload : JSON.stringify(payload, null, 2);
    outputEl.style.display = "block";
    outputEl.textContent = `${title}\n\n${printable}`;
}

function downloadText(filename, text, mime) {
    const blob = new Blob([text], { type: mime || "text/plain" });
    const url = URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.href = url;
    a.download = filename;
    document.body.appendChild(a);
    a.click();
    a.remove();
    URL.revokeObjectURL(url);
}
