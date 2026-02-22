#lang racket

(define (main)
  (displayln (+ (read-integer) (read-integer))))

(define (read-integer)
  (define m (regexp-match #rx"[0-9]+" (read-line)))
  (string->number (first m)))

(module+ main
  (main))

(provide main)
