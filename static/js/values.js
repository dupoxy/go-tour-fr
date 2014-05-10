/* Copyright 2012 The Go Authors.   All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */
'use strict';

angular.module('tour.values', []).

// List of modules with description and lessons in it.
value('TableOfContents', [{
    'id': 'mechanics',
    'title': 'Utilisation du tour',
    'description': '<p>Bienvenue au tour du <a href="http://golang.org">langage de programmation Go</a>. Ce tour couvre les caractéristiques les plus importantes du language, principalement:</p>',
    'lessons': ['welcome']
}, {
    'id': 'basics',
    'title': 'Principes de base',
    'description': '<p>Le point de départ, apprendre tous les rudiments du langage</p>Déclaration des variables, des fonctions d\'appel, et toutes les choses que vous devez savoir avant de passer aux prochaines leçons.</p>',
    'lessons': ['basics', 'flowcontrol', 'moretypes']
}, {
    'id': 'methods',
    'title': 'Methods et interfaces',
    'description': '<p>Savoir comment définir des méthodes sur les types, comment déclarer des interfaces, et comment assembler le tout.</p>',
    'lessons': ['methods']
}, {
    'id': 'concurrency',
    'title': 'Concurrency',
    'description': '<p>Go offre des primitives de simultanéité dans le cadre du langage de base.<p></p>Ce module les présentent et donne quelques exemples sur la façon de les utilisés pour mettre en œuvre différents modèles de concurrence.</p>',
    'lessons': ['concurrency']
}]).

// Translation
value('Translation', {
    "off": "Désactivé",
    "on": "Activé",
    "syntax": "Syntaxe",
    "lineno": "Numéros de ligne",
    "reset": "Réinitialiser",
    "format": "Formater",
    "kill": "Tuer",
    "run": "Exécuter",
    "compile": "Compiler et Exécuter",
    "more": "Options",
    "toc": "Menu",
    "prev": "Précédent",
    "next": "Suivant",
    "waiting": "En attente du serveur distant...",
    "errcomm": "Erreur de communication avec le serveur distant.",
}).

// Config for codemirror plugin
value('ui.config', {
    codemirror: {
        mode: 'text/x-go',
        matchBrackets: true,
        lineNumbers: true,
        autofocus: true,
        indentWithTabs: true,
        lineWrapping: true,
        extraKeys: {
            "Shift-Enter": function() {
                $('#run').click();
            },
            "PageDown": function() {
                return false;
            },
            "PageUp": function() {
                return false;
            },
            "Shift-Space": function() {
                $('#format').click();
            },
        }
    }
});
