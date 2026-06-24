/**
 * IM Music - Persistent Player Overlay
 * Menjaga audio tetap berjalan saat berpindah halaman
 * dengan overlay full-screen + mini player bar di bawah.
 */
(function () {
    'use strict';

    const API_BASE = 'http://localhost:8080';

    // localStorage keys
    const SK = {
        songId:      'im_songId',
        songTitle:   'im_songTitle',
        songArtist:  'im_songArtist',
        coverUrl:    'im_coverUrl',
        audioUrl:    'im_audioUrl',
        currentTime: 'im_currentTime',
        isPlaying:   'im_isPlaying',
    };

    let audio       = null;   // shared audio element
    let isPlaying   = false;
    let overlayOpen = false;

    /* ──────────────────────────────────────────────────────────
       AUDIO ENGINE
    ────────────────────────────────────────────────────────── */

    function getAudio() {
        if (audio) return audio;
        audio = new Audio();

        audio.addEventListener('timeupdate', () => {
            localStorage.setItem(SK.currentTime, audio.currentTime);
            if (!audio.duration) return;
            const pct = (audio.currentTime / audio.duration) * 100;
            setEl('po-fill',    el => el.style.width   = pct + '%');
            setEl('po-current', el => el.innerText = fmt(audio.currentTime));
        });

        audio.addEventListener('loadedmetadata', () => {
            setEl('po-total', el => el.innerText = fmt(audio.duration));
        });

        audio.addEventListener('ended', () => {
            isPlaying = false;
            localStorage.setItem(SK.isPlaying, 'false');
            setPlayIcons('play');
        });

        return audio;
    }

    /* ──────────────────────────────────────────────────────────
       LOAD & PLAY
    ────────────────────────────────────────────────────────── */

    async function loadAndPlay(songId) {
        const token = localStorage.getItem('token');
        if (!token) return;

        try {
            const res = await fetch(`${API_BASE}/api/songs/${songId}`, {
                headers: { Authorization: `Bearer ${token}` }
            });
            if (!res.ok) return;
            const song = (await res.json()).data;

            // Simpan state ke localStorage
            localStorage.setItem(SK.songId,      song.ID);
            localStorage.setItem(SK.songTitle,   song.Title);
            localStorage.setItem(SK.songArtist,  song.Artist);
            localStorage.setItem(SK.coverUrl,    song.CoverURL  || '');
            localStorage.setItem(SK.audioUrl,    song.AudioURL  || '');
            localStorage.setItem(SK.currentTime, '0');
            localStorage.setItem(SK.isPlaying,   'true');

            renderOverlayUI(song);
            renderMiniUI(song);

            const a = getAudio();
            if (song.AudioURL) {
                a.src = API_BASE + song.AudioURL;
                a.currentTime = 0;
                a.play()
                    .then(() => { isPlaying = true; setPlayIcons('pause'); })
                    .catch(console.warn);
            }

            openOverlay();
            showMiniBar();
            saveHistory(parseInt(songId), token);

        } catch (e) { console.error('PlayerOverlay:', e); }
    }

    /* ──────────────────────────────────────────────────────────
       RESTORE STATE (saat pindah halaman)
    ────────────────────────────────────────────────────────── */

    function restoreState() {
        const id        = localStorage.getItem(SK.songId);
        const url       = localStorage.getItem(SK.audioUrl);
        const wasPlaying= localStorage.getItem(SK.isPlaying) === 'true';
        const savedTime = parseFloat(localStorage.getItem(SK.currentTime) || '0');

        if (!id || !url) return;

        const song = {
            ID:       id,
            Title:    localStorage.getItem(SK.songTitle)  || '-',
            Artist:   localStorage.getItem(SK.songArtist) || '-',
            CoverURL: localStorage.getItem(SK.coverUrl)   || '',
            AudioURL: url,
        };

        renderOverlayUI(song);
        renderMiniUI(song);
        showMiniBar();

        const a = getAudio();
        a.src = API_BASE + url;

        a.addEventListener('canplay', function handler() {
            a.currentTime = savedTime;
            if (wasPlaying) {
                a.play()
                    .then(() => { isPlaying = true; setPlayIcons('pause'); })
                    .catch(() => setPlayIcons('play')); // autoplay blocked → user harus klik play
            } else {
                setPlayIcons('play');
            }
            a.removeEventListener('canplay', handler);
        });
    }

    /* ──────────────────────────────────────────────────────────
       INJECT HTML
    ────────────────────────────────────────────────────────── */

    function injectOverlay() {
        if (document.getElementById('playerOverlay')) return;
        const el = document.createElement('div');
        el.id = 'playerOverlay';
        el.innerHTML = `
            <header class="player-header">
                <a href="#" id="po-close"><i class="fas fa-chevron-down"></i></a>
                <span class="player-label">Sedang Diputar</span>
                <a href="#"><i class="fas fa-ellipsis-v"></i></a>
            </header>
            <main class="player-content">
                <div class="player-cover" id="po-cover">
                    <i class="fas fa-music" id="po-cover-icon"></i>
                </div>
                <div class="player-info">
                    <h3 id="po-title">-</h3>
                    <p id="po-artist">-</p>
                </div>
                <div class="player-progress">
                    <div class="progress-bar-track" id="po-track">
                        <div class="progress-bar-fill" id="po-fill"></div>
                    </div>
                    <div class="progress-time">
                        <span id="po-current">0:00</span>
                        <span id="po-total">0:00</span>
                    </div>
                </div>
                <div class="player-controls">
                    <i class="fas fa-random control-secondary"></i>
                    <i class="fas fa-step-backward"></i>
                    <div class="control-play" id="po-play-btn">
                        <i class="fas fa-play" id="po-play-icon"></i>
                    </div>
                    <i class="fas fa-step-forward"></i>
                    <i class="fas fa-redo control-secondary"></i>
                </div>
            </main>
        `;

        const container = document.querySelector('.mobile-container');
        if (container) container.appendChild(el);

        el.querySelector('#po-close').addEventListener('click', e => {
            e.preventDefault();
            PlayerOverlay.hide();
        });
        el.querySelector('#po-play-btn').addEventListener('click', () => PlayerOverlay.toggle());
        el.querySelector('#po-track').addEventListener('click', e => PlayerOverlay.seek(e));
    }

    function injectMiniBar() {
        if (document.getElementById('miniPlayerBar')) return;
        const bar = document.createElement('div');
        bar.id = 'miniPlayerBar';
        bar.innerHTML = `
            <div class="mini-left" id="mini-expand">
                <div class="mini-cover" id="mini-cover"></div>
                <div class="mini-text">
                    <span class="mini-title"  id="mini-title">-</span>
                    <span class="mini-artist" id="mini-artist">-</span>
                </div>
            </div>
            <div class="mini-play-btn" id="mini-play-btn">
                <i class="fas fa-play" id="mini-play-icon"></i>
            </div>
        `;

        const container = document.querySelector('.mobile-container');
        const nav = container && container.querySelector('.bottom-nav');
        if (nav) container.insertBefore(bar, nav);
        else if (container) container.appendChild(bar);

        bar.querySelector('#mini-expand').addEventListener('click', () => PlayerOverlay.show());
        bar.querySelector('#mini-play-btn').addEventListener('click', e => {
            e.stopPropagation();
            PlayerOverlay.toggle();
        });
    }

    /* ──────────────────────────────────────────────────────────
       RENDER UI
    ────────────────────────────────────────────────────────── */

    function renderOverlayUI(song) {
        const cover  = document.getElementById('po-cover');
        const icon   = document.getElementById('po-cover-icon');
        const title  = document.getElementById('po-title');
        const artist = document.getElementById('po-artist');
        if (!title) return;

        title.innerText  = song.Title  || '-';
        artist.innerText = song.Artist || '-';

        if (song.CoverURL) {
            cover.style.backgroundImage = `url('${API_BASE}${song.CoverURL}')`;
            if (icon) icon.style.display = 'none';
        } else {
            cover.style.backgroundImage = '';
            if (icon) icon.style.display = '';
        }
        // Reset progress when new song
        setEl('po-fill',    el => el.style.width = '0%');
        setEl('po-current', el => el.innerText   = '0:00');
        setEl('po-total',   el => el.innerText   = '0:00');
    }

    function renderMiniUI(song) {
        const cover  = document.getElementById('mini-cover');
        const title  = document.getElementById('mini-title');
        const artist = document.getElementById('mini-artist');
        if (!title) return;

        title.innerText  = song.Title  || '-';
        artist.innerText = song.Artist || '-';
        if (cover) {
            cover.style.backgroundImage = song.CoverURL
                ? `url('${API_BASE}${song.CoverURL}')` : '';
        }
    }

    /* ──────────────────────────────────────────────────────────
       OVERLAY SHOW / HIDE
    ────────────────────────────────────────────────────────── */

    function openOverlay() {
        const o = document.getElementById('playerOverlay');
        if (o) { o.classList.add('active'); overlayOpen = true; }
        const m = document.getElementById('miniPlayerBar');
        if (m) m.style.display = 'none';
    }

    function closeOverlay() {
        const o = document.getElementById('playerOverlay');
        if (o) { o.classList.remove('active'); overlayOpen = false; }
        showMiniBar();
    }

    function showMiniBar() {
        const m = document.getElementById('miniPlayerBar');
        if (m && !overlayOpen) {
            m.style.display = 'flex';
            // Beri ruang tambahan agar konten tidak tertutup mini player
            const mc = document.querySelector('.main-content');
            if (mc) mc.style.paddingBottom = '160px';
        }
    }

    /* ──────────────────────────────────────────────────────────
       HELPERS
    ────────────────────────────────────────────────────────── */

    function setEl(id, fn) {
        const el = document.getElementById(id);
        if (el) fn(el);
    }

    function setPlayIcons(state) {
        ['po-play-icon', 'mini-play-icon'].forEach(id => {
            const el = document.getElementById(id);
            if (!el) return;
            el.classList.toggle('fa-play',  state === 'play');
            el.classList.toggle('fa-pause', state === 'pause');
        });
    }

    function fmt(s) {
        if (!s || isNaN(s)) return '0:00';
        const m   = Math.floor(s / 60);
        const sec = Math.floor(s % 60);
        return `${m}:${sec.toString().padStart(2, '0')}`;
    }

    async function saveHistory(songId, token) {
        try {
            await fetch(`${API_BASE}/api/history`, {
                method: 'POST',
                headers: { Authorization: `Bearer ${token}`, 'Content-Type': 'application/json' },
                body: JSON.stringify({ song_id: songId })
            });
        } catch (e) {}
    }

    /* ──────────────────────────────────────────────────────────
       PUBLIC API
    ────────────────────────────────────────────────────────── */

    window.PlayerOverlay = {
        /** Muat dan putar lagu berdasarkan ID */
        play(songId) {
            loadAndPlay(songId);
        },

        /** Toggle play / pause */
        toggle() {
            const a = getAudio();
            if (!a.src) return;
            if (a.paused) {
                a.play().then(() => {
                    isPlaying = true;
                    localStorage.setItem(SK.isPlaying, 'true');
                    setPlayIcons('pause');
                }).catch(console.warn);
            } else {
                a.pause();
                isPlaying = false;
                localStorage.setItem(SK.isPlaying, 'false');
                setPlayIcons('play');
            }
        },

        /** Tampilkan overlay player */
        show()  { openOverlay();  },

        /** Sembunyikan overlay, kembali ke mini player */
        hide()  { closeOverlay(); },

        /** Seek audio berdasarkan klik pada progress bar */
        seek(event) {
            const a = getAudio();
            if (!a.duration) return;
            const track = document.getElementById('po-track');
            const rect  = track.getBoundingClientRect();
            const ratio = Math.max(0, Math.min(1, (event.clientX - rect.left) / rect.width));
            a.currentTime = ratio * a.duration;
        }
    };

    /* ──────────────────────────────────────────────────────────
       INIT
    ────────────────────────────────────────────────────────── */

    document.addEventListener('DOMContentLoaded', () => {
        injectOverlay();
        injectMiniBar();
        restoreState();
    });

})();
