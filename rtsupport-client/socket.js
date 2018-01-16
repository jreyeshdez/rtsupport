import {EventEmitter} from 'events';

export default class Socket {
    constructor(ws = new WebSocket(), ee = new EventEmitter()) {
        this.ws = ws;
        this.ee = ee;
        ws.onmessage = this.message.bind(this);
        ws.onopen = this.open.bind(this);
        ws.onclose = this.close.bind(this);
    }

    on(name, fn) {
        this.ee.on(name, fn);
    }

    off(name, fn) {
        this.ee.removeListener(name, fn);
    }

    emit(name, data) {
        const message = JSON.stringify({name, data});
        this.ws.send(message);
    }
    
    message(e) {
        try {
            const message = JSON.parse(e.data);
            this.ee.emit(message.name, message.data);
        } catch (error) {
            this.ee.emit('error', error);    
        }
    }

    open(e) {
        this.ee.emit('connect');
    }

    close(e) {
        this.ee.emit('disconnect');
    }
}