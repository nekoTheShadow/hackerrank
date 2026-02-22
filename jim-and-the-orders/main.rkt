#lang racket

(provide main)

(struct customer (name order time))

(define (main)
  (define n (string->number (read-line)))
  (define customers (readlines n))
  (displayln (string-join (map (lambda (c) (number->string (customer-name c)))
                               (sort customers cmp-customer))
                          " "))
  (void))

(define (readlines n)
  (for/list ([name (in-range n)])
    (let ([line (map string->number (string-split (read-line)))])
      (customer (+ name 1) (first line) (second line)))))

(define (cmp-customer c1 c2)
  (let ([t1 (+ (customer-time c1) (customer-order c1))]
        [t2 (+ (customer-time c2) (customer-order c2))])
    (if (= t1 t2)
        (< (customer-name c1) (customer-name c2))
        (< t1 t2))))

(module+ main
  (main))
