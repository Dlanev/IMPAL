/**
 * IM Music - Shared Actions (playlist + favourite)
 * Exposes window.Actions with:
 *   Actions.addToPlaylist(songId) -> opens a bottom sheet to pick / create a playlist
 *   Actions.like(songId)          -> adds the song to favourites
 *   Actions.toast(msg)            -> small transient notification
 * Works inside the iframe page context used by player-overlay.js.
 */
(function () {
    'use strict';

    const API_BASE = 'http://localhost:8080';

    function token() {
        return localStorage.getItem('token');
    }

    function authHeaders(extra) {
        return Object.assign(
            { 'Authorization': 'Bearer ' + token() },
            extra || {}
        );
    }

    function escapeHtml(str) {
        return String(str == null ? '' : str)
            .replace(/&/g, '&amp;')
            .replace(/</g, '&lt;')
            .replace(/>/g, '&gt;')
            .replace(/"/g, '&quot;')
            .replace(/'/g, '&#39;');
    }

    /* ── styling (injected once) ─────────────────────────────── */
    function injectStyles() {
        if (document.getElementById('actions-style')) return;
        const s = document.createElement('style');
        s.id = 'actions-style';
        s.textContent = `
        .am-backdrop{position:fixed;inset:0;background:rgba(0,0,0,.6);display:flex;align-items:flex-end;justify-content:center;z-index:9999;}
        .am-sheet{background:#181818;width:100%;max-width:480px;border-radius:16px 16px 0 0;padding:16px 16px 24px;max-height:70%;overflow-y:auto;font-family:'Montserrat',sans-serif;}
        .am-sheet h3{color:#fff;margin:0 0 12px;font-size:16px;}
        .am-item{display:flex;align-items:center;gap:12px;padding:12px 8px;color:#fff;cursor:pointer;border-radius:8px;font-size:14px;}
        .am-item:hover{background:#282828;}
        .am-item i{color:#1DB954;width:20px;text-align:center;}
        .am-empty{color:#b3b3b3;font-size:14px;padding:8px;}
        .am-close{background:#333;color:#fff;border:none;border-radius:20px;padding:10px 16px;width:100%;margin-top:8px;cursor:pointer;font-size:14px;font-family:'Montserrat',sans-serif;}
        .am-toast{position:fixed;bottom:150px;left:50%;transform:translateX(-50%);background:#1DB954;color:#fff;padding:10px 18px;border-radius:20px;z-index:10000;font-family:'Montserrat',sans-serif;font-size:14px;box-shadow:0 4px 12px rgba(0,0,0,.4);}
        .song-actions{display:flex;gap:14px;align-items:center;}
        .song-actions i{color:#b3b3b3;cursor:pointer;font-size:16px;padding:4px;}
        .song-actions i:hover{color:#fff;}
        `;
        document.head.appendChild(s);
    }

    function toast(msg) {
        injectStyles();
        const t = document.createElement('div');
        t.className = 'am-toast';
        t.innerText = msg;
        document.body.appendChild(t);
        setTimeout(() => { t.remove(); }, 2200);
    }

    function closeSheet() {
        const b = document.querySelector('.am-backdrop');
        if (b) b.remove();
        setMiniBarHidden(false);
    }

    // The mini player bar is injected by player-overlay.js into the TOP
    // window and floats over this page's iframe, so it would paint on top
    // of our bottom sheet. Temporarily hide it while a sheet is open.
    function topDoc() {
        try {
            return (window.top && window.top.document)
                ? window.top.document : document;
        } catch (e) { return document; }
    }

    function setMiniBarHidden(hidden) {
        const m = topDoc().getElementById('miniPlayerBar');
        if (m) m.style.visibility = hidden ? 'hidden' : '';
    }

    /* ── API calls ───────────────────────────────────────────── */
    async function fetchPlaylists() {
        const res = await fetch(`${API_BASE}/api/playlists`, {
            headers: authHeaders()
        });
        if (!res.ok) return [];
        const j = await res.json();
        return j.data || [];
    }

    // Create a playlist then return its id (backend does not return it,
    // so we re-fetch and match by name; list is ordered newest-first).
    async function createPlaylistGetId(name) {
        const res = await fetch(`${API_BASE}/api/playlists`, {
            method: 'POST',
            headers: authHeaders({ 'Content-Type': 'application/json' }),
            body: JSON.stringify({ name: name, description: '' })
        });
        if (!res.ok) return null;
        const list = await fetchPlaylists();
        const match = list.find(p => p.name === name);
        return match ? match.id : null;
    }

    async function addSongToPlaylist(playlistId, songId) {
        try {
            const res = await fetch(
                `${API_BASE}/api/playlists/${playlistId}/songs/${songId}`,
                { method: 'POST', headers: authHeaders() }
            );
            const j = await res.json().catch(() => ({}));
            toast(res.ok ? 'Ditambahkan ke playlist' : (j.message || 'Gagal menambahkan'));
        } catch (e) {
            console.error(e);
            toast('Kesalahan jaringan');
        }
    }

    async function like(songId) {
        if (!token()) { window.location.href = 'login.html'; return; }
        try {
            const res = await fetch(`${API_BASE}/api/favorites/${songId}`, {
                method: 'POST',
                headers: authHeaders()
            });
            const j = await res.json().catch(() => ({}));
            toast(res.ok ? 'Ditambahkan ke Disukai' : (j.message || 'Gagal menyukai'));
        } catch (e) {
            console.error(e);
            toast('Kesalahan jaringan');
        }
    }

    /* ── add-to-playlist bottom sheet ────────────────────────── */
    async function addToPlaylist(songId) {
        if (!token()) { window.location.href = 'login.html'; return; }
        injectStyles();
        closeSheet();
        setMiniBarHidden(true);

        const backdrop = document.createElement('div');
        backdrop.className = 'am-backdrop';
        backdrop.addEventListener('click', e => {
            if (e.target === backdrop) closeSheet();
        });

        const sheet = document.createElement('div');
        sheet.className = 'am-sheet';
        sheet.innerHTML = `<h3>Tambah ke Playlist</h3><div class="am-empty">Memuat playlist...</div>`;
        backdrop.appendChild(sheet);
        document.body.appendChild(backdrop);

        const playlists = await fetchPlaylists();

        let html = '<h3>Tambah ke Playlist</h3>';
        html += `<div class="am-item" data-action="new"><i class="fas fa-plus"></i><span>Buat playlist baru</span></div>`;
        if (playlists.length === 0) {
            html += `<div class="am-empty">Belum ada playlist. Buat satu di atas.</div>`;
        } else {
            playlists.forEach(p => {
                html += `<div class="am-item" data-id="${p.id}"><i class="fas fa-list-ul"></i><span>${escapeHtml(p.name)}</span></div>`;
            });
        }
        html += `<button class="am-close">Batal</button>`;
        sheet.innerHTML = html;

        sheet.querySelector('.am-close').addEventListener('click', closeSheet);

        sheet.querySelectorAll('.am-item').forEach(item => {
            item.addEventListener('click', async () => {
                if (item.dataset.action === 'new') {
                    const name = prompt('Nama playlist baru:');
                    if (!name || !name.trim()) return;
                    const pid = await createPlaylistGetId(name.trim());
                    if (pid != null) {
                        await addSongToPlaylist(pid, songId);
                    } else {
                        toast('Gagal membuat playlist');
                    }
                    closeSheet();
                } else {
                    await addSongToPlaylist(item.dataset.id, songId);
                    closeSheet();
                }
            });
        });
    }

    /* ── public API ──────────────────────────────────────────── */
    window.Actions = {
        addToPlaylist,
        like,
        toast,
        fetchPlaylists,
        createPlaylistGetId,
        addSongToPlaylist,
        API_BASE
    };

})();
